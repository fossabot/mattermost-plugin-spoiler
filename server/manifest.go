// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.github.moussetc.mattermost.plugin.spoiler",
  "name": "Spoiler Command",
  "description": "This plugin defines a /spoiler command.",
  "version": "2.1.4",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "SpoilerMode",
        "display_name": "Display mode for spoiler messages",
        "type": "radio",
        "help_text": "This setting will not affect native apps (Android,...).",
        "placeholder": "",
        "default": "button",
        "options": [
          {
            "display_name": "Spoiler button",
            "value": "button"
          },
          {
            "display_name": "Highlight ('redacted' look)",
            "value": "redacted"
          }
        ]
      },
      {
        "key": "IntegrationURL",
        "display_name": "Integration URL",
        "type": "text",
        "help_text": "Internal URL used by the MM server to call the plugin. Only use if using the SiteURL MM setting is not possible (for example if you get 509 HTTP errors), leave empty otherwise)",
        "placeholder": "http://localhost:8065",
        "default": ""
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
