package users

import (
	"github.com/dmazzella--/GoBasha_users-api/domain/users"
	"github.com/dmazzella--/GoBasha_users-api/services"
	"github.com/dmazzella--/GoBasha_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body ")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, createErr := services.CreateUser(user)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SearchUser")
}
