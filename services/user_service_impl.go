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
