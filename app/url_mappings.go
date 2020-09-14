package app

import (
	"github.com/dmazzella--/GoBasha_users-api/controllers/ping"
	"github.com/dmazzella--/GoBasha_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users/create", users.CreateUser)
	router.POST("/users/update/:user_id", users.CreateUser)

	router.GET("/users/search", users.SearchUser)
	router.GET("/users/get/:user_id", users.GetUser)
}
