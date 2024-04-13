package userRepo

import (
	"log"

	"github.com/upper/db/v4"
)

type UserRepo struct {
	Users db.Collection
}

func New(sess db.Session) UserRepo {
	coll := sess.Collection("user")

	return UserRepo{
		Users: coll,
	}
}

func (repo UserRepo) GetAll() []User {
	users := []User{
		{
			UserData: UserData{ID: "elo", Name: "elo"},
			Email:    "",
		},
	}

	err := repo.Users.Find().All(&users)
	if err != nil {
		log.Fatal("Błąd:", err)
	}
	return users
}

func (repo UserRepo) CreateUser(user UserCreate) (*UserData, error) {
	record, err := repo.Users.Insert(user)
	if err != nil {
		return nil, err
	}

	newUser := User{
		UserData: UserData{
			ID:   record.ID().(string),
			Name: user.Name,
		},
		Email: user.Email,
	}

	return &newUser.UserData, nil
}
