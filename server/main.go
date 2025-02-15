package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/overal-x/formatio/config"
	"github.com/overal-x/formatio/handlers"
	"github.com/overal-x/formatio/models"
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

	githubHandler := handlers.NewGithubHandler(db)

	srv := config.NewServer()

	srv.GET("/api/providers/github/", githubHandler.CreateApp)
	srv.GET("/api/providers/github/apps/", githubHandler.ListApps)
	srv.POST("/api/deploy/github/", githubHandler.DeployRepo)

	srv.Logger.Fatal(srv.Start(fmt.Sprintf(":%d", env.PORT)))
}
