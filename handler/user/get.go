package user

import (
	"log"

	. "framework.learning/handler"
	"framework.learning/model"
	"framework.learning/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	log.Print("getting user: " + username)
	user, err := model.GetUser(username)
	if err != nil {
		log.Print(err)
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
