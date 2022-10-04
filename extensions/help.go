package extensions

import (
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"
)

// HelpCommand covers the case where a user doesn't specify any command for help.
// Commands register their own help commands.
var HelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) == 1 {
			bot.Reply(
				message, "My trigger is: ! (eg: !help). Commands: help, uptime, boof, "+
					"forest, seen. Specify !help <command> to get more information",
			)

			return true
		}

		return false
	},
}
