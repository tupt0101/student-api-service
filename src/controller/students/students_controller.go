package students

import (
	"github.com/gin-gonic/gin"
	"github.com/tupt0101/student-api-service/src/oauth"
)

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
}

func Create(c *gin.Context) {

}

func GetById(c *gin.Context) {

}

func Update(c *gin.Context) {

}
