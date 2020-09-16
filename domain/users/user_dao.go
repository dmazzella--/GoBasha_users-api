package users

import (
	"fmt"
	"github.com/dmazzella--/GoBasha_users-api/datasources/mysql/users_db"
	"github.com/dmazzella--/GoBasha_users-api/utils/date_utils"
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail  = "unique_email"
	noRowsInResultSet = "no rows in result set"
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?)"
	queryGetUser      = "SELECT id, first_name, last_name, email, date_created from users where id = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noRowsInResultSet) {
			return errors.NewNotFoundError(fmt.Sprintf("user with id %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d :%s", user.Id, err.Error()))
	}
	// return nil error
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	user.Id = userId
	return nil
}
