package extensions

import (
	"fmt"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"
)

var BoofCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!boof")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, fmt.Sprintf("%s, it's free if you boof it", message.From))

		return true
	},
}

var BoofHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help boof")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, "POV: the bot is your dealer. Usage: !boof")

		return true
	},
}
