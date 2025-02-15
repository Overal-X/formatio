package services

import (
	"net/url"
	"reflect"

	"github.com/overal-x/formatio/types"
)

type IGithubService interface {
	CreateApp(args types.CreateAppArgs) (*types.CreateAppResult, error)
	InstallApp() error
	UpdateAppAccess() error
	UpdateAppWebhook() error
	CloneRepo() error
}

type GithubService struct{}

type KeyPair struct {
	KeyID      string
	PrivateKey string
}

func (g *GithubService) GeneratePrivateKey(appId int64, token string) (*KeyPair, error) {
	// ctx := context.Background()

	// Create GitHub client with token
	// ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	// tc := oauth2.NewClient(ctx, ts)
	// client := github.NewClient(tc)

	// Generate new private key using GitHub API
	// key, _, err := client.Apps.CreatePrivateKey(ctx, appId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to generate private key: %w", err)
	// }

	return &KeyPair{
		KeyID:      "key.GetKeyID()",
		PrivateKey: "key.GetKey()",
	}, nil
}

func (g *GithubService) CreateApp(args types.CreateAppArgs) (*types.CreateAppResult, error) {
	baseURL := "https://github.com/settings/apps/new"

	// Create URL values
	params := url.Values{}

	// Use reflection to iterate over struct fields
	v := reflect.ValueOf(args)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get the json tag
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}

		// Add non-empty values to params
		switch value.Kind() {
		case reflect.String:
			if str := value.String(); str != "" {
				params.Add(tag, str)
			}
		case reflect.Bool:
			if value.Bool() {
				params.Add(tag, "true")
			}
		}
	}

	// Build final URL
	if len(params) > 0 {
		baseURL += "?" + params.Encode()
	}

	return &types.CreateAppResult{
		AppUrl:     baseURL,
		PrivateKey: "",
	}, nil
}

func (g *GithubService) InstallApp() error {
	panic("unimplemented")
}

func (g *GithubService) UpdateAppAccess() error {
	panic("unimplemented")
}

func (g *GithubService) UpdateAppWebhook() error {
	panic("unimplemented")
}

func (g *GithubService) CloneRepo() error {
	panic("unimplemented")
}

func NewGithubService() IGithubService {
	return &GithubService{}
}
