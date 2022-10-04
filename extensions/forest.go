package extensions

import (
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"
)

// ForestCommand provides functionality for Electric Forest
var ForestCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!forest")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) == 1 {
			bot.Reply(message, message.From+", registration and dates. Usage: !forest registration")

			return true
		}

		if len(parts) > 1 && parts[1] == "registration" {
			bot.Reply(message, message.From+", registration for Electric Forest 2023 is not yet scheduled.")

			return true
		}

		if len(parts) > 1 && parts[1] == "dates" {
			bot.Reply(message, message.From+", dates for Electric Forest 2023 are not set yet.")

			return true
		}

		return true
	},
}

var ForestHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help forest")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) == 2 {
			bot.Reply(message, "registration, dates")

			return true
		}

		if len(parts) == 3 && parts[2] == "registration" {
			bot.Reply(message, "Registration date and link. Usage: !forest registration")

			return true
		}

		if len(parts) == 3 && parts[2] == "dates" {
			bot.Reply(message, "Dates of the next Electric Forest. Usage: !forest dates")

			return true
		}

		return true
	},
}
