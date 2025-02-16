package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/overal-x/formatio/config"
	"github.com/overal-x/formatio/handlers"
	"github.com/overal-x/formatio/models"
	"github.com/overal-x/formatio/services"
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

	err = db.AutoMigrate(&models.GithubApp{})
	if err != nil {
		log.Fatal(err)
	}

	githubService := services.NewGithubService()
	githubHandler := handlers.NewGithubHandler(db, githubService)

	srv := config.NewServer()

	srv.GET("/api/github/", githubHandler.CreateApp)
	srv.GET("/api/github/apps/", githubHandler.ListApps)
	srv.GET("/api/github/installations/:appId/", githubHandler.ListInstallations)
	srv.GET("/api/github/repos/:installationId/", githubHandler.ListRepo)
	srv.POST("/api/github/deploy/", githubHandler.DeployRepo)

	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", env.PORT)))
}
