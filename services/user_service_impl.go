package services

import (
	"context"
	"encoding/json"
	"go-rest-api/models"

	"github.com/google/uuid"
	"github.com/syndtr/goleveldb/leveldb"
)

type UserServiceImpl struct {
	userCollection *leveldb.DB
	ctx            context.Context
}

func NewUserService(userCollection *leveldb.DB, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (usi *UserServiceImpl) CreateUser(user *models.User) error {
	user.PersonalCode = uuid.New()

	pc, err := json.Marshal(&user.PersonalCode)
	u, err := json.Marshal(&user)

	usi.userCollection.Put([]byte(pc), []byte(u), nil)
	// func createUser(db *leveldb.DB) func(ctx *gin.Context) {
	// 	return func(ctx *gin.Context) {
	// 		newUser := User{}
	// 		err := ctx.ShouldBindJSON(&newUser)
	// 		if err != nil {
	// 			ctx.AbortWithError(http.StatusBadRequest, err)
	// 			return
	// 		}
	// 		newUser.PersonalCode = uuid.New()
	// 		pc, err := json.Marshal(&newUser.PersonalCode)
	// 		u, err := json.Marshal(&newUser)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		db.Put([]byte(pc), []byte(u), nil)
	// 		ctx.JSON(http.StatusCreated, &newUser)
	// 		return
	// 	}
	// }
	return err
}

func (usi *UserServiceImpl) GetUser(personalCode string) (*models.User, error) {
	var user models.User
	pc, err := json.Marshal(personalCode)
	if err != nil {
		panic(err)
	}
	value, err := usi.userCollection.Get([]byte(pc), nil)
	json.Unmarshal([]byte(value), &user)
	return &user, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(personalCode *string) error {
	return nil
}
