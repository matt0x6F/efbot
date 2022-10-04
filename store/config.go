package store

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var AppConfig *ApplicationConfig

type ApplicationConfig struct {
	botConfig *BotConfig
	mutex     sync.RWMutex
}

func (a *ApplicationConfig) DBLocation() string {
	return a.botConfig.DB.Location
}

// AllChannels returns a slice of all channels
func (a *ApplicationConfig) AllChannels() []string {
	var channels []string

	for _, channel := range a.botConfig.Channels {
		channels = append(channels, channel.Name)
	}

	return channels
}

// TrackedChannels returns a slice of tracked channels
func (a *ApplicationConfig) TrackedChannels() []string {
	var channels []string

	for _, channel := range a.botConfig.Channels {
		if channel.Track {
			channels = append(channels, channel.Name)
		}
	}

	return channels
}

// IsTracked indicates whether a channel should be tracked
func (a *ApplicationConfig) IsTracked(query string) bool {
	for _, channel := range a.botConfig.Channels {
		if channel.Track && channel.Name == query {
			return true
		}

		if !channel.Track && channel.Name == query {
			break
		}
	}

	return false
}

// IsTriggerChannel indicates whether a channel allows triggerable events
func (a *ApplicationConfig) IsTriggerChannel(query string) bool {
	for _, channel := range a.botConfig.Channels {
		if channel.Trigger && channel.Name == query {
			return true
		}

		if !channel.Trigger && channel.Name == query {
			break
		}
	}

	return false
}

// IsMod indicates whether a hostname correlates to an admin
func (a *ApplicationConfig) IsMod(query string) bool {
	for _, host := range a.botConfig.ModHosts {
		if host == query {
			return true
		}
	}

	return false
}

// IsModChannel indicates whether the channel can have moderator commands run in it
func (a *ApplicationConfig) IsModChannel(query string) bool {
	for _, channel := range a.botConfig.Channels {
		if channel.Name == query && channel.Moderator {
			return true
		}

		if channel.Name == query && !channel.Moderator {
			break
		}
	}

	return false
}

// ModeratedChannels returns a slice of moderated channels
func (a *ApplicationConfig) ModeratedChannels() []string {
	var channels []string

	for _, channel := range a.botConfig.Channels {
		if channel.Moderate {
			channels = append(channels, channel.Name)
		}
	}

	return channels
}

// ContactEmail returns the contact email field
func (a *ApplicationConfig) ContactEmail() string {
	return a.botConfig.ContactEmail
}

// RealName returns the real name field for IRC
func (a *ApplicationConfig) RealName() string {
	return a.botConfig.IRC.RealName
}

// Nick returns the nick field for IRC
func (a *ApplicationConfig) Nick() string {
	return a.botConfig.IRC.Nick
}

// Password returns the password field for IRC
func (a *ApplicationConfig) Password() string {
	return a.botConfig.IRC.Password
}

type ChannelConfig struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Trigger  bool   `yaml:"trigger"`
	Moderate bool   `yaml:"moderate"`
	// indicates the channel is one that can have moderator commands run in it
	Moderator bool `yaml:"moderator"`
	// indicates whether to track messages and users in a channel
	Track bool `yaml:"track"`
}

type BotConfig struct {
	// Used during Nickserv registration
	ContactEmail string `yaml:"contact_email"`
	IRC          struct {
		Nick string `yaml:"nick"`
		// Set on IRC user
		RealName string `yaml:"real_name"`
		// Used as the Nickserv password
		Password string `yaml:"password"`
	} `yaml:"irc"`
	ModHosts []string        `yaml:"mod_hosts"`
	Channels []ChannelConfig `yaml:"channels"`
	DB       struct {
		Location string `yaml:"location"`
	} `yaml:"db"`
}

func LoadAppConfig(location string) error {
	var botconfig = new(BotConfig)

	data, err := os.ReadFile(location)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, botconfig)

	AppConfig = new(ApplicationConfig)
	AppConfig.botConfig = botconfig

	return err
}
