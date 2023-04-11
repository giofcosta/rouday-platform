package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/giofcosta/rouday-api/cmd/api/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WeekConfigurationRepository interface {
	GetByID(id string) (*models.WeekConfiguration, error)
	Create(config *models.WeekConfiguration) error
	Update(config *models.WeekConfiguration) error
	UpdateRoutine(configId string, routine *models.Routine) error
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
		log.Println(err)
		return nil, err
	}

	var config models.WeekConfiguration

	collection := r.db.Database("rouday").Collection("weekConfigurations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&config)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &config, nil
}

func (r *weekConfigurationRepository) Create(config *models.WeekConfiguration) error {
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

func (r *weekConfigurationRepository) Update(config *models.WeekConfiguration) error {
	collection := r.db.Database("rouday").Collection("weekConfigurations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": config.ID}
	update := bson.M{
		"$set": bson.M{
			"week_start":       config.WeekStart,
			"work_hours_per_day": config.WorkHoursPerDay,
			"util_days_per_week":   config.UtilDaysPerWeek,
			"routines": config.Routines,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = fmt.Errorf("no document with id %s was found", config.ID)
		log.Println(err.Error())
		return err
	}

	return nil
}

func (r *weekConfigurationRepository) UpdateRoutine(configId string, routine *models.Routine) error {
	collection := r.db.Database("rouday").Collection("weekConfigurations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": configId, "routines.id": routine.ID}
	update := bson.M{"$set": bson.M{"routines.$": routine}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		err = fmt.Errorf("no document with id %s was found", routine.ID)
		log.Println(err.Error())
		return err
	}

	return nil
}









