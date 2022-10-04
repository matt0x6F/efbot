package extensions

import (
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"
)

// Listens for whois data. Some events may trigger a whois, this listener receives the whois commands and sets them
// appropriately in the DB
var WhoIsListener = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == "311" || message.Command == "319" || message.Command == "312" || message.
			Command == "671" || message.Command == "330" || message.Command == "318"
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Logger.Debug("WhoIs Listener", "command", message.Command, "content", message.Content, "params",
			message.Params)

		// todo: whois data can be aggregated by this point

		return true
	},
}

var WhoIsCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return strings.HasPrefix(message.Content, "!whois")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) == 2 {
			bot.Send("WHOIS " + parts[1])
		} else {
			bot.Send("WHOIS " + message.From)
		}

		return true
	},
}

var WhoIsHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help whois")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, "Performs a whois on the given target. "+
			"Usage: !whois <user>")

		return false
	},
}
