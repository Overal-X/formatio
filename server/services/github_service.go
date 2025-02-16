package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v69/github"
)

type IGithubService interface {
	CloneRepo() error
	GetAppToken(GetAppTokenArgs) (*string, error)
	GetInstallationToken(GetInstallationTokenArgs) (*string, error)
}

type GithubService struct{}

func (g *GithubService) CloneRepo() error {
	panic("unimplemented")
}

type GetAppTokenArgs struct {
	ClientId   string
	PrivateKey string
}

func (g *GithubService) GetAppToken(args GetAppTokenArgs) (*string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = jwt.MapClaims{
		"iss": args.ClientId,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(args.PrivateKey))
	if err != nil {
		return nil, err
	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

type GetInstallationTokenArgs struct {
	ClientId       string
	PrivateKey     string
	InstallationId int64
}

func (g *GithubService) GetInstallationToken(args GetInstallationTokenArgs) (*string, error) {
	token, err := g.GetAppToken(GetAppTokenArgs{
		ClientId:   args.ClientId,
		PrivateKey: args.PrivateKey,
	})
	if err != nil {
		return nil, err
	}

	client := github.NewClient(nil).WithAuthToken(*token)
	installationToken, _, err := client.Apps.CreateInstallationToken(context.Background(), args.InstallationId, nil)
	if err != nil {
		return nil, err
	}

	return installationToken.Token, nil
}

func NewGithubService() IGithubService {
	return &GithubService{}
}
