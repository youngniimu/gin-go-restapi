package main

import (
	"context"
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
	ctx            context.Context
	userCollection *leveldb.DB
	err            error
)

func init() {
	ctx = context.TODO()
	userCollection, err = leveldb.OpenFile("db", nil)
	if err != nil {
		log.Fatal(err)
	}
	userservice = services.NewUserService(userCollection, ctx)
	UserController = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	basepath := server.Group("/v1")
	UserController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run())
}
