package main

import (
	"go-rest-api/controllers"
	"go-rest-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	UserController controllers.UserController
	userCollection *leveldb.DB
	err            error
)

func init() {
	userCollection, err = leveldb.OpenFile("db", nil)
	if err != nil {
		log.Fatal(err)
	}
	userservice = services.NewUserService(userCollection)
	UserController = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	basepath := server.Group("/user")
	UserController.RegisterUserRoutes(basepath)
	log.Fatal(server.Run())
}
