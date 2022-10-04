package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// User is used by pop to map your users database table to your go code.
type User struct {
	gorm.Model
	Host string
	User string
	Nick string
}

// String is not required by pop and may be deleted
func (u *User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// KnownUser is a rolling list of users which the bot knows through various means
type KnownUser struct {
	gorm.Model
	Host    string
	User    string
	Nick    string
	Event   string
	Channel string
}

// String is not required by pop and may be deleted
func (u *KnownUser) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}
