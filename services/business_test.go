package services

import (
	"go-locate/model"
	"go-locate/repository"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	db, err := model.InitDB(os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
	if err != nil {
		log.Fatalf("Error instantiating database %v\n", err)
	}

	businessRepo := repository.NewBusiness(db)
	businessService := NewBusiness(businessRepo)

	user := &model.User{
		Username: "testuser",
		Email:    "testuser@gmail.com",
		Password: "password",
	}

	arg := &model.Business{
		Name:  "Test Business",
		Email: "test@gmail.com",
		PhoneNumbers: []model.Contact{
			{
				Contact: "237679165995",
			},
		},
		Categories: []model.BusinessCategory{
			{
				CategoryId: 1,
			},
		},
		Description: "Description for test business",
		Location: []model.BusinessLocation{
			{
				MainLocation: "Buea",
				Description:  "Test description",
				Map:          "This is supposed to be an Iframe of a map",
			},
		},
	}

	business, e := businessService.Create(arg.Name, arg.Description, arg.Email, arg.Location, user, arg.PhoneNumbers, arg.Categories)

	require.NoError(t, e)
	require.NotEmpty(t, business)
}
