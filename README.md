# Disclaimer

**This fork builds on the original Mattermost WelcomeBot plugin, allowing channel moderators to manage team and channel welcome messages and actions directly through Mattermost commands — no server file modifications required.**

## Configuration

1. Go to **System Console > Plugins > Management** and click **Enable** to enable the Welcome Bot plugin.
    - If you are running Mattermost v5.11 or earlier, you must first go to the [releases page of this GitHub repository](https://github.com/mattermost/mattermost-plugin-welcomebot/releases), download the latest release, and upload it to your Mattermost instance [following this documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).

2. Please ensure that you go to **System Console > Site Configuration > Users and Teams** and verify the setting for **Enable users to open Direct Message channels with:**. If the value of this field is set to `Any member of the team`, you'll need to add the welcome bot manually to all the teams where it needs to be included.

## Usage
You can preview configured messages and create a channel welcome message using the following bot commands:  

### Team Default Channels Management
These commands allow you to manage channels that users automatically join when entering the current team.

- `/welcomebot add_team_default_channels <[~channel]>`  
  Add channels to automatically join when entering the current team.
- `/welcomebot remove_team_default_channels <[~channel]>`  
  Remove channels from the auto-join list when entering the current team.

### Channel Welcome Messages  
These commands manage welcome messages for **the current channel**.

- `/welcomebot set_personal_channel_welcome_message [welcome-message]`  
  Set a personal welcome message for the current channel (Direct channels are not supported).
- `/welcomebot get_personal_channel_welcome_message`  
  Display the personal welcome message set for the current channel (if any).
- `/welcomebot delete_personal_channel_welcome_message`  
  Delete the personal welcome message for the current channel (if any).
- `/welcomebot set_published_channel_welcome_message [welcome-message]`  
  Set a published welcome message for the current channel (Direct channels are not supported).
- `/welcomebot get_published_channel_welcome_message`  
  Display the published welcome message set for the current channel (if any).
- `/welcomebot delete_published_channel_welcome_message`  
  Delete the published welcome message for the current channel (if any).
- `/welcomebot list_channel_welcomes`  
  List all channels with configured welcome messages.

### Team Welcome Messages  
These commands manage the welcome message displayed after joining **the current team**.

- `/welcomebot set_team_welcome_message [message]`  
  Set the welcome message displayed after joining the current team.
- `/welcomebot get_team_welcome_settings`  
  Display the welcome settings set for the current team.
- `/welcomebot delete_team_welcome_message`  
  Delete the welcome message for the current team.

### Message Visibility  

- **Personal Channel Welcome Messages**  
  Sent as direct messages to the user and also posted in the channel, but only visible to the user. Note that channel messages visible only to a specific user are ephemeral, meaning they may disappear quickly, even if the user hasn't read them.
- **Published Channel Welcome Messages**  
  Sent to the channel and visible to all channel members.  
- **Team Welcome Messages**  
  Sent as direct messages to the user.  

### Message Templates  

Welcome messages support a simple templating system for user-specific placeholders:  

- `{{.UserDisplayName}}` → Displays the user's full name (e.g., `John Doe`).  
- `{{.UserHandleName}}` → Displays the user's handle (e.g., `@john_doe`).  

When previewing a message, all template placeholders will be replaced with the current user's details.  


## Development

This plugin contains a server portion. Read our documentation about the [Developer Workflow](https://developers.mattermost.com/integrate/plugins/developer-workflow/) and [Developer Setup](https://developers.mattermost.com/integrate/plugins/developer-setup/) for more information about developing and extending plugins.

### Releasing new versions

The version of a plugin is determined at compile time, automatically populating a `version` field in the [plugin manifest](plugin.json):
* If the current commit matches a tag, the version will match after stripping any leading `v`, e.g. `1.3.1`.
* Otherwise, the version will combine the nearest tag with `git rev-parse --short HEAD`, e.g. `1.3.1+d06e53e1`.
* If there is no version tag, an empty version will be combined with the short hash, e.g. `0.0.0+76081421`.

To disable this behaviour, manually populate and maintain the `version` field.
