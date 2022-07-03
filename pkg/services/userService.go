package services

import (
	"fmt"

	"github.com/svetlimladenov/collegesystems/pkg/models"
)

type UserService struct {
	test string
}

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Email     string
}

func (us UserService) Login(name, password string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	var user User
	db.First(&user)

	fmt.Printf("us.test: %v\n", us.test)
	us.test = name
	fmt.Println(us, name, password)
	fmt.Printf("us.test: %v\n", us.test)
	return nil
}

func (us UserService) Register(model models.RegisterInputModel) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	var user User
	result := db.Where("Username = ?", model.Username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
