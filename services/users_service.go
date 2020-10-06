package services

import (
	"github.com/dmazzella--/GoBasha_users-api/domain/users"
	"github.com/dmazzella--/GoBasha_users-api/utils/crypto"
	"github.com/dmazzella--/GoBasha_users-api/utils/date_utils"
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User, bool) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusCreated
	user.Password = crypto.GetMd5(user.Password, user.Email)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(user users.User, isPatch bool) (*users.User, *errors.RestErr) {
	current := &users.User{Id: user.Id}

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
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	current := &users.User{Id: userId}

	if err := current.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
