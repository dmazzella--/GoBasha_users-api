package users

import (
	"fmt"
	"github.com/dmazzella--/GoBasha_users-api/datasources/mysql/users_db"
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, status, date_created, password) VALUES (?,?,?,?,?,?)"
	queryGetUser      = "SELxECT id, first_name, last_name, email, date_created, status from users where id = ?"
	queryUpdateUser   = "UPDATE users set first_name = ?, last_name = ?, email = ? , status = ? where id = ?"
	queryDeleteUser   = "DELETE FROM users where id = ?"
	queryFindByStatus = "SELECT id, first_name, last_name, email, date_created, status from users where status = ?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E001GetUser, err)
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {

		return errors.LogAndNewInternalServerError(errors.E002GetUser, err)
	}
	// return nil no errors
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E003SaveUser, err)
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.DateCreated, user.Password)
	if saveErr != nil {
		return errors.LogAndNewInternalServerError(errors.E004SaveUser, err)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E005SaveUser, err)
	}
	user.Id = userId

	// return nil no errors
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E006UpdateUser, err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E007UpdateUser, err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E008DeleteUser, err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return errors.LogAndNewInternalServerError(errors.E009DeleteUser, err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.LogAndNewInternalServerError(errors.E010FindByStatus, err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.LogAndNewInternalServerError(errors.E011FindByStatus, err)
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.LogAndNewInternalServerError(errors.E012FindByStatus, err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
