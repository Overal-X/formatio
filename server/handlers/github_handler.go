package handlers

import (
	"context"
	"strconv"

	"github.com/google/go-github/v69/github"
	"github.com/labstack/echo/v4"
	"github.com/overal-x/formatio/models"
	"github.com/overal-x/formatio/services"
	"github.com/overal-x/formatio/utils"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type IGithubHandler interface {
	CreateApp(c echo.Context) error
	ListApps(c echo.Context) error
	DeployRepo(c echo.Context) error
	ListInstallations(c echo.Context) error
	ListRepo(c echo.Context) error
}

type GithubHandler struct {
	db            *gorm.DB
	githubService services.IGithubService
}

// @ID create-app
// @Success 201 {object} models.GithubApp
// @Router /api/providers/github [get]
func (g *GithubHandler) CreateApp(c echo.Context) error {
	// authId := c.QueryParam("authId")
	next := c.QueryParam("next")
	code := c.QueryParam("code")

	client := github.NewClient(nil)
	ghApp, resp, err := client.Apps.CompleteAppManifest(context.Background(), code)
	if err != nil {
		return c.JSON(resp.StatusCode, echo.Map{"message": err})
	}

	app := models.GithubApp{
		AppId:         ghApp.GetID(),
		AppName:       ghApp.GetName(),
		OwnerId:       ghApp.GetOwner().GetID(),
		OwnerUsername: ghApp.GetOwner().GetLogin(),
		OwnerType:     ghApp.GetOwner().GetType(),
		ClientId:      ghApp.GetClientID(),
		ClientSecret:  ghApp.GetClientSecret(),
		WebhookSecret: ghApp.GetWebhookSecret(),
		PrivateKey:    ghApp.GetPEM(),
	}
	err = g.db.Create(&app).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	if next != "" {
		return c.Redirect(302, next)
	}

	return c.JSON(201, app)
}

// @ID list-apps
// @Success 200 {array} models.GithubApp
// @Router /api/providers/github/apps [get]
func (g *GithubHandler) ListApps(c echo.Context) error {
	apps := []models.GithubApp{}
	err := g.db.Find(&apps).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	return c.JSON(200, apps)
}

func (g *GithubHandler) DeployRepo(c echo.Context) error {
	apps := []models.GithubApp{}
	err := g.db.Find(&apps).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	return c.JSON(200, apps)
}

func (g *GithubHandler) ListInstallations(c echo.Context) error {
	app := models.GithubApp{}
	err := g.db.First(&app).Where("id = ?", c.Param("appId")).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	appToken, err := g.githubService.GetAppToken(services.GetAppTokenArgs{
		ClientId:   app.ClientId,
		PrivateKey: app.PrivateKey,
	})
	if err != nil {
		return err
	}

	client := github.NewClient(nil).WithAuthToken(*appToken)
	installations, _, err := client.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		return err
	}

	return c.JSON(200, installations)
}

func (g *GithubHandler) ListRepo(c echo.Context) error {
	app := models.GithubApp{}
	err := g.db.First(&app, c.Param("id")).Error
	if err != nil {
		return utils.HandleGormError(c, err)
	}

	appToken, err := g.githubService.GetInstallationToken(services.GetInstallationTokenArgs{
		InstallationId: int64(lo.Must(strconv.Atoi(c.Param("installationId")))),
		ClientId:       app.ClientId,
		PrivateKey:     app.PrivateKey,
	})
	if err != nil {
		return err
	}

	client := github.NewClient(nil).WithAuthToken(*appToken)
	installations, _, err := client.Apps.ListRepos(context.Background(), &github.ListOptions{})
	if err != nil {
		return err
	}

	return c.JSON(200, installations)
}

func NewGithubHandler(db *gorm.DB, githubService services.IGithubService) IGithubHandler {
	return &GithubHandler{db: db, githubService: githubService}
}
