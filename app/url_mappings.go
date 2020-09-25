package app

import (
	"github.com/dmazzella--/GoBasha_users-api/controllers/ping"
	"github.com/dmazzella--/GoBasha_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// CRUD
	// C
	router.POST("/users/create", users.Create)

	// R
	router.GET("/users/search", users.Search)
	router.GET("/users/get/:user_id", users.Get)

	//R^i
	router.GET("/internal/users/search/", users.Search)

	// U
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)

	//D
	router.DELETE("/users/:user_id", users.Delete)

}
