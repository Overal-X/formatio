package types

type Installation struct {
	Id            *int64   `json:"id,omitempty"`
	AppId         int64    `json:"app_id,omitempty"`
	AppName       string   `json:"app_name,omitempty"`
	AppSlug       string   `json:"app_slug,omitempty"`
	OwnerId       int64    `json:"owner_id"`
	OwnerUsername string   `json:"owner_username"`
	OwnerType     string   `json:"owner_type"`
	Events        []string `json:"events,omitempty"`
}

type Repo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	OwnerId       int64  `json:"owner_id"`
	OwnerUsername string `json:"owner_username"`
	OwnerType     string `json:"owner_type"`
}
