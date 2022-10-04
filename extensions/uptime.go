package extensions

import (
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"
)

var UptimeCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!uptime")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, bot.Uptime())

		return true
	},
}

var UptimeHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help uptime")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) > 1 && parts[1] == "uptime" {
			bot.Reply(message, "Specifies the start time and the amount of time the bot has been online. "+
				"Usage: !uptime")

			return true
		}

		return false
	},
}
