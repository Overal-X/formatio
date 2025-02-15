package models

import "time"

type GithubApp struct {
	Id        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	AppId         int64  `json:"appId"`
	AppName       string `json:"appName"`
	OwnerId       int64  `json:"ownerId"`
	OwnerUsername string `json:"ownerUsername"`
	OwnerType     string `json:"ownerType"`
	ClientId      string `json:"-"`
	ClientSecret  string `json:"-"`
	WebhookSecret string `json:"-"`
	PrivateKey    string `json:"-"`
}
