package services

import (
	"github.com/dmazzella--/GoBasha_users-api/domain/users"
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUser(user users.User, isPatch bool) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPatch {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	current, err := GetUser(userId)
	if err != nil {
		return err
	}
	if err := current.Delete(); err != nil {
		return err
	}
	return nil
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
