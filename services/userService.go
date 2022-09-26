package services

import (
	"fmt"

	"github.com/svetlimladenov/collegesystems/models"
	"github.com/svetlimladenov/collegesystems/pkg/hash"
)

type UserService struct{}

func (us UserService) Login(name, password string) error {
	db, err := GetDB() // move to db pkg :)
	if err != nil {
		return err
	}

	var user models.User
	db.First(&user)
	return nil
}

func (us UserService) Register(user models.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	result := db.Where("Username = ?", user.Username).First(&models.User{})
	if result.Error == nil {
		return fmt.Errorf("user already registered")
	}

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.RoleId = 1

	db.Create(&user)

	return nil
}
