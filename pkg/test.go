package pkg

import (
	"go-locate/model"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func SetupTest(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	err = model.InitDB(os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
	if err != nil {
		t.Fatalf(err.Error())
	}
	model.ClearTables(&model.User{}, &model.Business{}, &model.BusinessCategory{}, &model.Category{}, &model.Contact{})
}

func ShutdownTest(cleanup ...func()) {
	for _, c := range cleanup {
		c()
	}
	model.CloseDB()
}
