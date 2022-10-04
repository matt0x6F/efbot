package extensions

import (
	"errors"
	"fmt"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
	"gopkg.in/sorcix/irc.v2"
	"gorm.io/gorm"

	"efbot/models"
	"efbot/store"
)

var Track = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return store.AppConfig.IsTracked(message.To) && message.To != store.AppConfig.Nick() && message.
			From != store.AppConfig.Nick() && message.From != "ChanServ"
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		go LogMessage(bot, message)

		return false
	},
}

var LastSeenCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && store.AppConfig.
			IsTriggerChannel(message.To) && strings.HasPrefix(message.Content, "!seen")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		parts := strings.Split(message.Content, " ")

		if len(parts) <= 1 {
			return false
		}

		user := parts[1]

		if parts[1] == message.From {
			bot.Reply(message, fmt.Sprintf("%s, %s is online now", message.From, user))

			return true
		}

		msg := models.Message{}
		err := store.DB.Transaction(func(tx *gorm.DB) error {
			result := tx.Where("nick = ?", user).Last(&msg)

			return result.Error
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			bot.Reply(message, fmt.Sprintf("%s, I haven't seen %s", message.From, user))

			return true
		}
		if err != nil {
			bot.Logger.Error("Failed to find message by user in message log", "err", err)

			return true
		}

		// this logic relies on the bot observing all quits and parts, which is not totally realistic. it may get out of
		// sync from time to time.
		if msg.Event != irc.PART && msg.Event != irc.QUIT {
			bot.Reply(message, fmt.Sprintf("%s, %s is online now", message.From, user))
		} else {
			bot.Reply(message, fmt.Sprintf("%s, last seen %s in %s at %s", message.From, msg.Nick, msg.To,
				msg.UpdatedAt.String()))
		}

		return true
	},
}

var LastSeenHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help seen")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, "describes when and where a user was last seen. Usage: !seen <nickname>")

		return true
	},
}

var WhoAmICommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!whoami")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		if store.AppConfig.IsMod(message.Host) {
			bot.Reply(message, fmt.Sprintf("%s, you are a mod", message.From))
		} else {
			bot.Reply(message, fmt.Sprintf("%s, never heard of her", message.From))
		}

		return true
	},
}

var WhoAmIHelpCommand = hbot.Trigger{
	Condition: func(bot *hbot.Bot, message *hbot.Message) bool {
		return message.Command == irc.PRIVMSG && strings.HasPrefix(message.Content, "!help seen")
	},
	Action: func(bot *hbot.Bot, message *hbot.Message) bool {
		bot.Reply(message, "indicates who the person is to a bot. Usage: !whoami")

		return true
	},
}

// LogMessage makes a best effort to log a message in the database. It will report errors,
// but errors are not meant to disrupt application flow.
func LogMessage(bot *hbot.Bot, message *hbot.Message) {
	err := store.DB.Transaction(func(tx *gorm.DB) error {
		msg := models.Message{
			Host:      message.Host,
			User:      message.User,
			Nick:      message.From,
			Content:   message.Content,
			To:        message.To,
			Timestamp: message.TimeStamp,
			Event:     message.Command,
		}

		result := tx.Create(&msg)

		return result.Error
	})
	if err != nil {
		bot.Logger.Error("Failed to insert message into message log", "err", err)
	}
}
