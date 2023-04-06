package repositories

import (
	"context"
	"log"
	"time"

	"github.com/giofcosta/rouday-api/cmd/api/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WeekConfigurationRepository interface {
	GetByID(id string) (*models.WeekConfiguration, error)
	CreateWeekConfiguration(config *models.WeekConfiguration) error
}

type weekConfigurationRepository struct {
	db *mongo.Client
}

func NewWeekConfigurationRepository(db *mongo.Client) WeekConfigurationRepository {
	return &weekConfigurationRepository{db}
}

func (r *weekConfigurationRepository) GetByID(id string) (*models.WeekConfiguration, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var config models.WeekConfiguration

	collection := r.db.Database("rouday").Collection("weekConfigurations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *weekConfigurationRepository) CreateWeekConfiguration(config *models.WeekConfiguration) error {
	collection := r.db.Database("rouday").Collection("weekConfigurations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, config)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}










