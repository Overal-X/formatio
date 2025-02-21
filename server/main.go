package main

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/overal-x/formatio/config"
	"github.com/overal-x/formatio/handlers"
	"github.com/overal-x/formatio/models"
	"github.com/overal-x/formatio/services"
	"github.com/overal-x/formatio/types"
	"github.com/samber/do"
	"gorm.io/gorm"
)

// @title Formatio API
// @version 1.0
// @BasePath /
func main() {
	i := do.New()

	do.Provide(i, config.NewEnv)
	do.Provide(i, config.NewDatabaseConnection)
	do.Provide(i, services.NewRabbitMQConnection)
	do.Provide(i, services.NewRabbitMQService)
	do.Provide(i, services.NewGithubService)
	do.Provide(i, services.NewExecService)
	do.Provide(i, services.NewNixpacksService)
	do.Provide(i, services.NewDeploymentService)
	do.Provide(i, services.NewFileService)
	do.Provide(i, services.NewProjectService)
	do.Provide(i, handlers.NewGithubHandler)
	do.Provide(i, handlers.NewProjectHandler)
	do.Provide(i, handlers.NewDeploymentHandler)

	env := do.MustInvoke[*config.Env](i)
	db := do.MustInvoke[*gorm.DB](i)

	err := db.AutoMigrate(
		&models.GithubApp{},
		&models.Project{},
		&models.Environment{},
		&models.Deployment{},
		&models.DeploymentLog{},
		&models.Network{},
	)
	if err != nil {
		log.Fatal(err)
	}

	rabbitmqService := do.MustInvoke[services.IRabbitMQService](i)

	githubHandler := do.MustInvoke[handlers.IGithubHandler](i)
	projectHandler := do.MustInvoke[handlers.IProjectHandler](i)
	deploymentHandler := do.MustInvoke[handlers.IDeploymentHandler](i)

	srv := config.NewServer()

	srv.GET("/api/github/", githubHandler.CreateApp)
	srv.GET("/api/github/apps/", githubHandler.ListApps)
	srv.GET("/api/github/installations/:app_id/", githubHandler.ListInstallations)
	srv.GET("/api/github/repos/:app_id/:installation_id/", githubHandler.ListRepo)
	srv.POST("/api/github/deploy/", githubHandler.DeployRepo)

	srv.GET("/api/projects/", projectHandler.List)
	srv.POST("/api/projects/", projectHandler.Create)
	srv.GET("/api/projects/:id/", projectHandler.Get)
	srv.PATCH("/api/projects/:id/", projectHandler.Update)
	srv.DELETE("/api/projects/:id/", projectHandler.Delete)
	srv.POST("/api/projects/:id/deploy/", projectHandler.Deploy)
	srv.GET("/api/projects/:id/network/", projectHandler.GetNetwork)

	srv.GET("/api/deployments/:project_id/", deploymentHandler.ListDeployments)
	srv.GET("/api/deployments/:deployment_id/logs/", deploymentHandler.ListDeploymentLogs)

	go func() {
		projectService := do.MustInvoke[services.IProjectService](i)

		err := rabbitmqService.Subscribe(services.SubscribeArgs{
			Queue: services.GITHUB_DEPLOYMENT_QUEUE,
			Callback: func(content string) error {
				payload := types.DeployArgs{}
				if err := json.Unmarshal([]byte(content), &payload); err != nil {
					return err
				}
				if err := projectService.HandleDeploy(payload); err != nil {
					return err
				}

				return nil
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}()

	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", env.PORT)))
	i.Shutdown()
}
