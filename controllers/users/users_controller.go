package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "GetUser "+c.Param("user_id"))
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "CreateUser")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SearchUser")
}
