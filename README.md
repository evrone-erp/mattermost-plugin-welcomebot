# Disclaimer

**This fork builds on the original Mattermost WelcomeBot plugin, allowing channel moderators to manage team and channel welcome messages and actions directly through Mattermost commands — no server file modifications required.**

# Welcome Bot Evrone Plugin

[![Build Status](https://img.shields.io/circleci/project/github/mattermost/mattermost-plugin-welcomebot/master)](https://circleci.com/gh/mattermost/mattermost-plugin-welcomebot)
[![Code Coverage](https://img.shields.io/codecov/c/github/mattermost/mattermost-plugin-welcomebot/master)](https://codecov.io/gh/mattermost/mattermost-plugin-welcomebot)
[![Release](https://img.shields.io/github/v/release/mattermost/mattermost-plugin-welcomebot)](https://github.com/mattermost/mattermost-plugin-welcomebot/releases/latest)
[![HW](https://img.shields.io/github/issues/mattermost/mattermost-plugin-welcomebot/Up%20For%20Grabs?color=dark%20green&label=Help%20Wanted)](https://github.com/mattermost/mattermost-plugin-welcomebot/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22Up+For+Grabs%22+label%3A%22Help+Wanted%22)

**Maintainer:** [@mickmister](https://github.com/mickmister)

Use this plugin to improve onboarding and HR processes. It adds a Welcome Bot that helps welcome users to teams and/or channels as well as easily join channels based on selections.

![image](https://user-images.githubusercontent.com/13119842/58736467-fd226400-83cb-11e9-827b-6bbe33d062ab.png)

*Welcome a new team member to Mattermost Contributors team. Then add the user to a set of channels based on their selection.*

![image](https://user-images.githubusercontent.com/13119842/58736540-30fd8980-83cc-11e9-8e8e-94ea3042b3b1.png)

## Configuration

1. Go to **System Console > Plugins > Management** and click **Enable** to enable the Welcome Bot plugin.
    - If you are running Mattermost v5.11 or earlier, you must first go to the [releases page of this GitHub repository](https://github.com/mattermost/mattermost-plugin-welcomebot/releases), download the latest release, and upload it to your Mattermost instance [following this documentation](https://docs.mattermost.com/administration/plugins.html#plugin-uploads).

2. Please ensure that you go to **System Console > Site Configuration > Users and Teams** and verify the setting for **Enable users to open Direct Message channels with:**. If the value of this field is set to `Any member of the team`, you'll need to add the welcome bot manually to all the teams where it needs to be included.

## Usage

The preview of the configured messages, as well as the creation of a channel welcome message, can be done via bot commands:
* `/welcomebot help` - Displays usage information.
* `/welcomebot list_channel_welcomes` - Lists channels with welcome messages
* `/welcomebot set_personal_channel_welcome [welcome-message]` - Sets the given text as current's channel personal welcome message.
* `/welcomebot get_personal_channel_welcome` - Gets the current channel's personal welcome message.
* `/welcomebot delete_personal_channel_welcome` - Deletes the current channel's personal welcome message.
* `/welcomebot set_published_channel_welcome [welcome-message]` - Sets the given text as current's channel published (visible for all) welcome message.
* `/welcomebot get_published_channel_welcome` - Gets the current channel's published (visible for all) welcome message.
* `/welcomebot delete_published_channel_welcome` - Deletes the current channel's published (visible for all) welcome message.

## Development

This plugin contains a server portion. Read our documentation about the [Developer Workflow](https://developers.mattermost.com/integrate/plugins/developer-workflow/) and [Developer Setup](https://developers.mattermost.com/integrate/plugins/developer-setup/) for more information about developing and extending plugins.

### Releasing new versions

The version of a plugin is determined at compile time, automatically populating a `version` field in the [plugin manifest](plugin.json):
* If the current commit matches a tag, the version will match after stripping any leading `v`, e.g. `1.3.1`.
* Otherwise, the version will combine the nearest tag with `git rev-parse --short HEAD`, e.g. `1.3.1+d06e53e1`.
* If there is no version tag, an empty version will be combined with the short hash, e.g. `0.0.0+76081421`.

To disable this behaviour, manually populate and maintain the `version` field.
