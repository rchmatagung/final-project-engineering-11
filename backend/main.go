package main

import (
	"github.com/rg-km/final-project-engineering-11/backend/config"
	"github.com/rg-km/final-project-engineering-11/backend/controller"
	"github.com/rg-km/final-project-engineering-11/backend/db/migration"
	"github.com/rg-km/final-project-engineering-11/backend/repository"
	"github.com/rg-km/final-project-engineering-11/backend/router"
	"github.com/rg-km/final-project-engineering-11/backend/service"
)

func main() {
	migration.Migrate()
	db := config.GetConnection()
	defer db.Close()
	router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
	router.Run(":8090")
}
