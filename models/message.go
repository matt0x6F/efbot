package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Message is used by pop to map your messages database table to your go code.
type Message struct {
	gorm.Model
	Timestamp time.Time
	Host      string
	User      string
	Nick      string
	Content   string
	To        string
	Event     string
}

// String is not required by pop and may be deleted
func (m *Message) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}
