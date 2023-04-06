package main

import (
	"github.com/giofcosta/rouday-api/cmd/api/data/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	WeekConfigurationRepository repositories.WeekConfigurationRepository
}

func NewApplication(dbClient *mongo.Client) *Application {
	weekConfigurationRepository := repositories.NewWeekConfigurationRepository(dbClient)
	return &Application{weekConfigurationRepository}
}	