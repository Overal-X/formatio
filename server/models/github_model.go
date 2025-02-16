package models

import "time"

type GithubApp struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	AppId         int64  `json:"app_id"`
	AppName       string `json:"app_name"`
	OwnerId       int64  `json:"owner_id"`
	OwnerUsername string `json:"owner_username"`
	OwnerType     string `json:"owner_type"`
	ClientId      string `json:"-"`
	ClientSecret  string `json:"-"`
	WebhookSecret string `json:"-"`
	PrivateKey    string `json:"-"`
}
