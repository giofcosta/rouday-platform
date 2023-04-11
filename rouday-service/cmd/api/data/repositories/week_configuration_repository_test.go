package repositories

import (
	"testing"
	"time"

	"github.com/giofcosta/rouday-api/cmd/api/data/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func createValidWeekConfiguration() *models.WeekConfiguration {
	id := uuid.New()
	layout := "2006-01-02T15:04:05Z"
	weekStart, _ := time.Parse(layout, "2023-01-02T00:00:00Z")
	return &models.WeekConfiguration{
		ID: id.String(),
		Routines: make([]models.Routine, 0),
		UtilDaysPerWeek: 5,
		WeekStart: weekStart,
		WorkHoursPerDay: 14,
	}
}

func TestCreateWeekConfiguration(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.RunOpts("create week configuration success", mtest.NewOptions(), func(mt *mtest.T) {
		repo := &weekConfigurationRepository{
			db: mt.Client,
		}
	
		weekConfig := createValidWeekConfiguration()
	
		mt.AddMockResponses(mtest.CreateSuccessResponse())
	
		err := repo.Create(weekConfig)
		if err != nil {
			t.Errorf("CreateWeekConfiguration() error: %v", err)
		}
	})
	

	mt.RunOpts("create week configuration duplicate", mtest.NewOptions(), func(mt *mtest.T) {
		repo := &weekConfigurationRepository{
			db: mt.Client,
		}

		weekConfig := createValidWeekConfiguration()
		_ = repo.Create(weekConfig)

		err := repo.Create(weekConfig)
		if err == nil {
			t.Error("CreateWeekConfiguration() should return an error when inserting a duplicate")
		}
	})

	mt.RunOpts("create week configuration invalid", mtest.NewOptions(), func(mt *mtest.T) {
		repo := &weekConfigurationRepository{
			db: mt.Client,
		}

		weekConfig := &models.WeekConfiguration{
		}

		err := repo.Create(weekConfig)
		if err == nil {
			t.Error("CreateWeekConfiguration() should return an error when inserting an invalid configuration")
		}
	})

	mt.RunOpts("create week configuration db error", mtest.NewOptions().ClientOptions(options.Client().SetWriteConcern(writeconcern.New(writeconcern.WMajority()))), func(mt *mtest.T) {
		repo := &weekConfigurationRepository{
			db: mt.Client,
		}

		mt.AddMockResponses(mtest.CreateCommandErrorResponse(
			mtest.CommandError{
				Code:    1000,
				Message: "database error",
				Name:    "DatabaseError",
			},
		))

		weekConfig := createValidWeekConfiguration()

		err := repo.Create(weekConfig)
		if err == nil {
			t.Error("CreateWeekConfiguration() should return an error when there is a database error")
		}
	})
}
