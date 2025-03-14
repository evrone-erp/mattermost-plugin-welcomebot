package core

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
)

func (p *Plugin) getSiteURL() string {
	siteURL := "http://localhost:8065"

	config := p.API.GetConfig()

	if config == nil || config.ServiceSettings.SiteURL == nil || len(*config.ServiceSettings.SiteURL) == 0 {
		return siteURL
	}

	return *config.ServiceSettings.SiteURL
}

func (p *Plugin) newSampleMessageTemplate(teamName string, userID string) (*MessageTemplate, error) {
	data := &MessageTemplate{}
	var err *model.AppError

	if data.User, err = p.API.GetUser(userID); err != nil {
		p.API.LogError("failed to query user", "user_id", userID, "err", err)
		return nil, fmt.Errorf("failed to query user %s: %w", userID, err)
	}

	if data.Team, err = p.API.GetTeamByName(strings.ToLower(teamName)); err != nil {
		p.API.LogError("failed to query team", "team_name", teamName, "err", err)
		return nil, fmt.Errorf("failed to query team %s: %w", teamName, err)
	}

	if data.Townsquare, err = p.API.GetChannelByName(data.Team.Id, "town-square", false); err != nil {
		p.API.LogError("failed to query town-square", "team_name", data.Team.Name)
		return nil, fmt.Errorf("failed to query town-square %s: %w", data.Team.Name, err)
	}

	if data.DirectMessage, err = p.API.GetDirectChannel(data.User.Id, p.BotUserID); err != nil {
		p.API.LogError("failed to query direct message channel", "user_name", data.User.Username)
		return nil, fmt.Errorf("failed to query direct message channel %s: %w", data.User.Id, err)
	}

	data.UserDisplayName = data.User.GetDisplayName(model.ShowNicknameFullName)

	return data, nil
}

func (p *Plugin) PreviewWelcomeMessage(teamName string, args *model.CommandArgs, configMessage ConfigMessage) error {
	messageTemplate, err := p.newSampleMessageTemplate(teamName, args.UserId)
	if err != nil {
		return err
	}

	post := p.renderWelcomeMessage(*messageTemplate, configMessage)
	post.ChannelId = args.ChannelId
	_ = p.API.SendEphemeralPost(args.UserId, post)

	return nil
}

func (p *Plugin) renderWelcomeMessage(messageTemplate MessageTemplate, configMessage ConfigMessage) *model.Post {
	actionButtons := make([]*model.PostAction, 0)

	for _, configAction := range configMessage.Actions {
		if configAction.ActionType == actionTypeAutomatic {
			action := &Action{}
			action.UserID = messageTemplate.User.Id
			action.Context = &ActionContext{}
			action.Context.TeamID = messageTemplate.Team.Id
			action.Context.UserID = messageTemplate.User.Id
			action.Context.Action = "automatic"

			for _, channelName := range configAction.ChannelsAddedTo {
				p.joinChannel(action, channelName)
			}
		}

		if configAction.ActionType == actionTypeButton {
			actionButton := &model.PostAction{
				Name: configAction.ActionDisplayName,
				Integration: &model.PostActionIntegration{
					Context: map[string]interface{}{
						"action":  configAction.ActionName,
						"team_id": messageTemplate.Team.Id,
						"user_id": messageTemplate.User.Id,
					},
					URL: fmt.Sprintf("%v/plugins/%v/addchannels", p.getSiteURL(), p.Manifest.Id),
				},
			}

			actionButtons = append(actionButtons, actionButton)
		}
	}

	tmpMsg, _ := template.New("Response").Parse(strings.Join(configMessage.Message, "\n"))
	var message bytes.Buffer
	err := tmpMsg.Execute(&message, messageTemplate)
	if err != nil {
		p.API.LogError(
			"Failed to execute message template",
			"err", err.Error(),
		)
	}

	post := &model.Post{
		Message: message.String(),
		UserId:  p.BotUserID,
	}

	if len(configMessage.AttachmentMessage) > 0 || len(actionButtons) > 0 {
		tmpAtch, _ := template.New("AttachmentResponse").Parse(strings.Join(configMessage.AttachmentMessage, "\n"))
		var attachMessage bytes.Buffer
		err := tmpAtch.Execute(&attachMessage, messageTemplate)
		if err != nil {
			p.API.LogError(
				"Failed to execute message template",
				"err", err.Error(),
			)
		}

		sa1 := &model.SlackAttachment{
			Text: attachMessage.String(),
		}

		if len(actionButtons) > 0 {
			sa1.Actions = actionButtons
		}

		attachments := make([]*model.SlackAttachment, 0)
		attachments = append(attachments, sa1)
		post.Props = map[string]interface{}{
			"attachments": attachments,
		}
	}

	return post
}

func (p *Plugin) processActionMessage(messageTemplate MessageTemplate, action *Action, configMessageAction ConfigMessageAction) {
	for _, channelName := range configMessageAction.ChannelsAddedTo {
		p.joinChannel(action, channelName)
	}

	tmpMsg, _ := template.New("Response").Parse(strings.Join(configMessageAction.ActionSuccessfulMessage, "\n"))
	var message bytes.Buffer
	err := tmpMsg.Execute(&message, messageTemplate)
	if err != nil {
		p.API.LogError(
			"Failed to execute message template",
			"err", err.Error(),
		)
	}

	post := &model.Post{
		Message:   message.String(),
		ChannelId: messageTemplate.DirectMessage.Id,
		UserId:    p.BotUserID,
	}

	if _, err := p.API.CreatePost(post); err != nil {
		p.API.LogError(
			"We could not create the response post",
			"user_id", post.UserId,
			"err", err.Error(),
		)
	}
}

func (p *Plugin) joinChannel(action *Action, channelName string) {
	if channel, err := p.API.GetChannelByName(action.Context.TeamID, channelName, false); err == nil {
		if _, err := p.API.AddChannelMember(channel.Id, action.Context.UserID); err != nil {
			p.API.LogError("Couldn't add user to the channel, continuing to next channel", "user_id", action.Context.UserID, "channel_id", channel.Id)
			return
		}
	} else {
		p.API.LogError("failed to get channel, continuing to the next channel", "channel_name", channelName, "user_id", action.Context.UserID)
	}
}
