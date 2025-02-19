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
)

// @title Formatio API
// @version 1.0
// @BasePath /
func main() {
	env := config.NewEnv()
	db, err := config.NewDatabaseConnection(env)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.GithubApp{}, &models.Project{}, &models.Environment{})
	if err != nil {
		log.Fatal(err)
	}

	githubService := services.NewGithubService()
	githubHandler := handlers.NewGithubHandler(db, githubService)

	execService := services.NewExecService()
	nixpacksService := services.NewNixpacksService(execService)

	rabbitmqConnection := services.NewRabbitMQConnection(env)
	rabbitmqService := services.NewRabbitMQService(rabbitmqConnection)

	projectService := services.NewProjectService(db, execService, nixpacksService, githubService, rabbitmqService)
	projectHandler := handlers.NewProjectHandler(projectService)

	srv := config.NewServer()

	srv.GET("/api/github/", githubHandler.CreateApp)
	srv.GET("/api/github/apps/", githubHandler.ListApps)
	srv.GET("/api/github/installations/:appId/", githubHandler.ListInstallations)
	srv.GET("/api/github/repos/:appId/:installationId/", githubHandler.ListRepo)
	srv.POST("/api/github/deploy/", githubHandler.DeployRepo)

	srv.GET("/api/projects/", projectHandler.List)
	srv.POST("/api/projects/", projectHandler.Create)
	srv.GET("/api/projects/:id/", projectHandler.Get)
	srv.PATCH("/api/projects/:id/", projectHandler.Update)
	srv.DELETE("/api/projects/:id/", projectHandler.Delete)
	srv.POST("/api/projects/:id/deploy/", projectHandler.Deploy)

	go rabbitmqService.Subscribe(services.SubscribeArgs{
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

	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", env.PORT)))
}
