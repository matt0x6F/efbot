package extensions

import (
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"

	"efbot/store"
)

// NickservCommand assists in getting started with Nickserv on Libera. It has workflows for registration,
// manual identification, verification, and cloaking. Logins are taken care of by Hellabot.
var NickservCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!nickserv")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) > 1 && parts[1] == "register" {
			bot.Msg("NickServ", "register "+bot.Password+" "+store.AppConfig.ContactEmail())

			bot.Reply(message, "Registered my nickname")

			return true
		}

		if len(parts) > 1 && parts[1] == "identify" {
			bot.Msg("NickServ", "identify "+bot.Password)

			bot.Reply(message, "Identified with Nickserv")

			return true
		}

		if len(parts) > 1 && parts[1] == "cloak" {
			bot.Join("#libera-cloak")
			bot.Msg("#libera-cloak", "!cloakme")

			bot.Reply(message, "Applied for cloak")

			return true
		}

		if len(parts) > 2 && parts[1] == "verify" {
			bot.Msg("NickServ", "verify register "+bot.Nick+" "+parts[2])

			bot.Reply(message, "Verified with Nickserv")

			return true
		}

		return false
	},
}

var NickservHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == "PRIVMSG" && strings.HasPrefix(message.Content, "!help nickserv")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) == 2 {
			bot.Reply(message, "cloak, identify, verify, register")

			return true
		}

		if len(parts) == 3 && parts[2] == "cloak" {
			bot.Reply(message, "Triggers the bot to apply for a cloak. Usage: !nickserv cloak")

			return true
		}

		if len(parts) == 3 && parts[2] == "identify" {
			bot.Reply(message, "Triggers a manual identification with Nickserv. Usage: !nickserv identify")

			return true
		}

		if len(parts) == 3 && parts[2] == "register" {
			bot.Reply(message, "Triggers the bot to register with Nickserv. Usage: !nickserv register")

			return true
		}

		if len(parts) == 3 && parts[2] == "verify" {
			bot.Reply(message, "Triggers the bot to verify with Nickserv. Usage: !nickserv verify <code>")

			return true
		}

		return true
	},
}
