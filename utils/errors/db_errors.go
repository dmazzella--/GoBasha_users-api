package errors

import "fmt"

type DBError struct {
	Detailed string
	Id       string
}

var (
	E001GetUser = DBError{
		"Error Trying to get a user",
		"E001",
	}

	E002GetUser = DBError{
		"error when trying to get user by id",
		"E002",
	}

	E003SaveUser = DBError{
		"error when trying to save a user",
		"E003",
	}

	E004SaveUser = DBError{
		"error when trying to save a user",
		"E004",
	}

	E005SaveUser = DBError{
		"error when trying to save a user",
		"E005",
	}

	E006UpdateUser = DBError{
		"error when trying to update a user",
		"E006",
	}

	E007UpdateUser = DBError{
		"error when trying to update a user",
		"E007",
	}

	E008DeleteUser = DBError{
		"error when trying to delete a user",
		"E008",
	}

	E009DeleteUser = DBError{
		"error when trying to delete a user",
		"E009",
	}

	E010FindByStatus = DBError{
		"error when trying to find a user",
		"E010",
	}

	E011FindByStatus = DBError{
		"error when trying to find a user",
		"E011",
	}

	E012FindByStatus = DBError{
		"error when trying to find a user",
		"E012",
	}
)

func (dberror *DBError) GetFormattedMessage() string {
	return fmt.Sprintf("%s - %s", dberror.Id, dberror.Detailed)
}

func (dberror *DBError) GetGenericMessage() string {
	return fmt.Sprintf("Internal Error with ID [%s]", dberror.Id)
}
