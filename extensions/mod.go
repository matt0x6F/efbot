package extensions

import (
	"fmt"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"

	"efbot/store"
)

var ModCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && store.AppConfig.IsMod(message.Host) && (store.AppConfig.
			IsModChannel(message.To) || store.AppConfig.Nick() == message.To) && strings.HasPrefix(message.
			Content, "!mod")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) <= 2 {
			return true
		}

		command := strings.ToLower(parts[1])

		switch command {
		case "warn":
			if len(parts) < 5 {
				bot.Reply(message, fmt.Sprintf("%s, the command structure is: !warn <#channel> <user> <reason>",
					message.From))

				return true
			}

			channel := parts[2]
			user := parts[3]
			reason := strings.Join(parts[4:], " ")

			if !strings.HasPrefix(channel, "#") {
				bot.Reply(message, fmt.Sprintf("%s, the command structure is: !warn <#channel> <user> <reason>. "+
					"Channels must start with a hash (#).",
					message.From))

				return true
			}

			bot.Msg(channel, fmt.Sprintf("%s, you have been warned. Reason: %s", user, reason))
			bot.Notice(message.From, "User warned successfully in "+channel)
		}

		return true
	},
}

var ModHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && store.AppConfig.IsMod(message.Host) && strings.HasPrefix(message.
			Content,
			"!help mod")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")
		if len(parts) == 2 {
			bot.Reply(message, "warn. Usage: !warn <#channel> <user> <reason>")

			return true
		}

		if parts[2] == "warn" {
			bot.Reply(message, "Warns a user on behalf of the moderator. "+
				"Reasons are required. Usage: !warn <#channel> <user> <reason>")

			return true
		}

		return true
	},
}
