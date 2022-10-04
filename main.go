package main

import (
	"log"

	hbot "github.com/whyrusleeping/hellabot"
	botlog "gopkg.in/inconshreveable/log15.v2"

	"efbot/extensions"
	"efbot/store"
)

func main() {
	err := store.LoadAppConfig("bot.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = store.Open()
	if err != nil {
		log.Fatal(err)
	}

	options := func(bot *hbot.Bot) {
		bot.Channels = store.AppConfig.AllChannels()
		bot.SASL = true
		bot.SSL = true
		bot.Realname = store.AppConfig.ContactEmail()
		bot.Password = store.AppConfig.Password()
	}

	bot, err := hbot.NewBot("irc.libera.chat:6697", store.AppConfig.Nick(), options)
	if err != nil {
		log.Fatal(err)
	}

	ogHandler := botlog.LvlFilterHandler(botlog.LvlDebug, botlog.StdoutHandler)
	bot.Logger.SetHandler(ogHandler)

	// Commands
	bot.AddTrigger(extensions.HelpCommand)
	// uptime
	bot.AddTrigger(extensions.UptimeCommand)
	bot.AddTrigger(extensions.UptimeHelpCommand)
	// nickserv
	bot.AddTrigger(extensions.NickservCommand)
	bot.AddTrigger(extensions.NickservHelpCommand)
	// forest
	bot.AddTrigger(extensions.ForestCommand)
	bot.AddTrigger(extensions.ForestHelpCommand)
	// boof
	bot.AddTrigger(extensions.BoofCommand)
	bot.AddTrigger(extensions.BoofHelpCommand)
	// about
	bot.AddTrigger(extensions.AboutCommand)
	bot.AddTrigger(extensions.AboutHelpCommand)
	// tracking
	bot.AddTrigger(extensions.Track)
	bot.AddTrigger(extensions.LastSeenCommand)
	bot.AddTrigger(extensions.LastSeenHelpCommand)
	// mod
	bot.AddTrigger(extensions.ModCommand)
	bot.AddTrigger(extensions.ModHelpCommand)
	// whoami
	bot.AddTrigger(extensions.WhoAmICommand)
	bot.AddTrigger(extensions.WhoAmIHelpCommand)
	// whois
	bot.AddTrigger(extensions.WhoIsListener)
	bot.AddTrigger(extensions.WhoIsCommand)
	bot.AddTrigger(extensions.WhoIsHelpCommand)

	bot.Run()
}
