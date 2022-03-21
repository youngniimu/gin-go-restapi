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
	value, err := usi.userCollection.Get([]byte(pc), nil)
	json.Unmarshal([]byte(value), &user)
	return &user, err
}

func (usi *UserServiceImpl) GetAll() ([]models.User, error) {
	var users []models.User
	var err error
	iter := usi.userCollection.NewIterator(nil, nil)
	for iter.Next() {
		var userInDb models.User
		err = json.Unmarshal([]byte(iter.Value()), &userInDb)
		users = append(users, userInDb)
	}
	return users, err
}

func (usi *UserServiceImpl) UpdateUser(user *models.User, personalCode string) error {
	pc, err := json.Marshal(personalCode)
	user.PersonalCode, err = uuid.Parse(personalCode)
	u, err := json.Marshal(&user)
	usi.userCollection.Put([]byte(pc), []byte(u), nil)
	return err
}

func (usi *UserServiceImpl) DeleteUser(personalCode string) error {
	pc, err := json.Marshal(personalCode)
	err = usi.userCollection.Delete([]byte(pc), nil)
	return err
}
