package main

import (
	"go-rest-api/models"
	"go-rest-api/services"
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestCreateUserAndGetUser(t *testing.T) {
	db, err := leveldb.OpenFile("_testdb", nil)
	if err != nil {
		t.Errorf("failed to open db")
	}
	testService := services.NewUserService(db)
	newUser := models.User{
		FirstName: "teemu",
		LastName:  "halme",
	}
	testService.CreateUser(&newUser)

	userInDb, _ := testService.GetUser(newUser.PersonalCode.String())

	if userInDb.FirstName != "teemu" || userInDb.LastName != "halme" {
		t.Errorf("user entry not saved correctly")
	}

	db.Close()
	os.Remove("_testdb")
}

