package app

import "bookstore_users_api/controllers"

func mapUrls(){
	router.GET("/ping", controllers.Ping)
}