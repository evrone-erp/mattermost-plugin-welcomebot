# Disclaimer

**This fork builds on the original Mattermost WelcomeBot plugin, allowing channel moderators to manage team and channel welcome messages and actions directly through Mattermost commands — no server file modifications required.**

## Configuration

1. Go to **System Console > Plugins > Management** and click **Enable** to enable the Welcome Bot plugin.
    - If you are running Mattermost v5.11 or earlier, you must first go to the [releases page of this GitHub repository](https://github.com/mattermost/mattermost-plugin-welcomebot/releases), download the latest release, and upload it to your Mattermost instance [following this documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).

2. Please ensure that you go to **System Console > Site Configuration > Users and Teams** and verify the setting for **Enable users to open Direct Message channels with:**. If the value of this field is set to `Any member of the team`, you'll need to add the welcome bot manually to all the teams where it needs to be included.

## Usage
You can preview configured messages and create a channel welcome message using the following bot commands:  

* `/welcomebot help` - Displays usage information.  
* `/welcomebot list_channel_welcomes` - Lists channels with configured welcome messages.  
* `/welcomebot set_personal_channel_welcome [welcome-message]` - Sets the given text as the current channel's personal welcome message.  
* `/welcomebot get_personal_channel_welcome` - Retrieves the current channel's personal welcome message.  
* `/welcomebot delete_personal_channel_welcome` - Deletes the current channel's personal welcome message.  
* `/welcomebot set_published_channel_welcome [welcome-message]` - Sets the given text as the current channel's published (visible to all members) welcome message.  
* `/welcomebot get_published_channel_welcome` - Retrieves the current channel's published (visible to all members) welcome message.  
* `/welcomebot delete_published_channel_welcome` - Deletes the current channel's published welcome message.  
* `/welcomebot get_team_welcome` - Displays the team welcome message and the default channels.  
* `/welcomebot set_team_welcome_message [welcome-message]` - Sets the current team's welcome message.  
* `/welcomebot remove_team_welcome` - Removes the current team's welcome message.  
* `/welcomebot add_team_default_channels [~channel_name,]` - Adds default channels for the team welcome message.  
* `/welcomebot remove_team_default_channels [~channel_name,]` - Removes default channels from the team welcome message.  

### Message Visibility  

- **Personal channel welcome messages** are sent as direct messages and also posted in the channel, but only visible to the user.  
- **Published channel welcome messages** are sent to the channel and visible to all channel members.  
- **Team welcome messages** are sent as direct messages.  

### Message Templates  

Welcome messages support a simple templating system:  

* `{{.UserDisplayName}}` - Example: `John Doe`  
* `{{.UserHandleName}}` - Example: `@john_doe`  

When previewing a message, all template placeholders will be replaced with the current user's details.  


## Development

This plugin contains a server portion. Read our documentation about the [Developer Workflow](https://developers.mattermost.com/integrate/plugins/developer-workflow/) and [Developer Setup](https://developers.mattermost.com/integrate/plugins/developer-setup/) for more information about developing and extending plugins.

### Releasing new versions

The version of a plugin is determined at compile time, automatically populating a `version` field in the [plugin manifest](plugin.json):
* If the current commit matches a tag, the version will match after stripping any leading `v`, e.g. `1.3.1`.
* Otherwise, the version will combine the nearest tag with `git rev-parse --short HEAD`, e.g. `1.3.1+d06e53e1`.
* If there is no version tag, an empty version will be combined with the short hash, e.g. `0.0.0+76081421`.

To disable this behaviour, manually populate and maintain the `version` field.
