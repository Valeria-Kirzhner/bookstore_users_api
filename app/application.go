package app

import "github.com/gin-gonic/gin"

var (
	// gin is creating a different go routine for each request. Thats why all my controllers need to be statless.
	router = gin.Default()
)
func StartApplication(){
	mapUrls()
	router.Run(":8080")
}