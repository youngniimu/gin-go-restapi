package main

import (
	"go-rest-api/models"
	"go-rest-api/services"
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestCreateUserAndGetUserAndUpdateUserAndDeleteUser(t *testing.T) {
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

	testService.UpdateUser(&models.User{FirstName: "tmu", LastName: "hlm"}, newUser.PersonalCode.String())

	userInDb, _ = testService.GetUser(newUser.PersonalCode.String())

	if userInDb.FirstName != "tmu" || userInDb.LastName != "hlm" {
		t.Errorf("user entry not updated correctly")
	}

	testService.DeleteUser(newUser.PersonalCode.String())

	allUsers, _ := testService.GetAll()

	if len(allUsers) != 0 {
		t.Errorf("user delete failed")
	}

	db.Close()
	os.RemoveAll("_testdb")
}
