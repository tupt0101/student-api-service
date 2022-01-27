package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()

	router.Run(":" + os.Getenv("APP_PORT"))
}
