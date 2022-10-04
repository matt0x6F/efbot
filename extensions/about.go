package extensions

import (
	"fmt"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"

	"efbot/store"
)

var AboutCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!about")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(
			message, fmt.Sprintf(
				"%s, %s is a custom bot. "+
					"It persists messages when indicated to by users and stores information related to commands.",
				message.From, store.AppConfig.Nick(),
			),
		)

		return true
	},
}

var AboutHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help about")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, "About the bot. Usage: !about")

		return true
	},
}
