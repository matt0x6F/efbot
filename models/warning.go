package models

import (
	"time"

	"gorm.io/gorm"
)

type Warning struct {
	gorm.Model
	Timestamp time.Time
	Host      string
	User      string
	Nick      string
	Reason    string
	// Hostname of the admin that authored the warning
	By string
}
