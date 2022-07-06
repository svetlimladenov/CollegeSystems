package services

import "github.com/svetlimladenov/collegesystems/models"

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

func (us UserService) Register(username, password string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	var user models.User
	result := db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
