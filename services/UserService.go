package services

import (
	"boilerplate-golang/models"
	"boilerplate-golang/utilities"
	"fmt"
)

func GetUsers() ([]*models.User, error) {

	var users []*models.User

	utilities.Conn.Find(&users)

	if len(users) <= 0 {
		return nil, fmt.Errorf("no records has found")
	}

	return users, nil
}

func GetUserByUsername(name *string) (*models.User, error) {
	var user *models.User

	utilities.Conn.Where("username LIKE ?", "%"+*name+"%").First(&user)

	if user.Username == "" {
		return nil, fmt.Errorf("unable to search query with name %s", *name)
	}

	return user, nil
}

func GetUserById(id *int) (*models.User, error) {
	var user *models.User

	utilities.Conn.First(&user, *id)

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to search query with id %d", *id)
	}

	return user, nil
}

func UpdateUserById(id *int) (*int, error) {
	return nil, nil
}

func DeleteUserById(id *int) (*int, error) {
	return nil, nil
}
