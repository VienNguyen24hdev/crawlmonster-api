package models

import (
	"time"

	"github.com/jasonbronson/crawlmonster-api/config"
)

func CreateUser(user *User) (bool, error) {
	result := config.Cfg.GormDB.Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func GetUserByEmail(email string) (*User, error) {
	user := User{}
	result := config.Cfg.GormDB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}

type User struct {
	ID                 string
	CustomerID         int
	Email              string
	Username           string
	Password           string
	Logins             int
	LastLogin          time.Time
	Active             int
	StripeActive       int
	StripeID           string
	StripeSubscription string
	StripePlan         string
	TrialEndsAt        time.Time
	SubscriptionEndsAt time.Time
}
