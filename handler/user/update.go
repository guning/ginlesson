package user

import (
	"log"
	"strconv"

	. "framework.learning/handler"
	"framework.learning/model"
	"framework.learning/pkg/errno"
	"framework.learning/util"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	log.Print("updating user: reqId: " + util.GetReqID(c))
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u.Id = uint64(userId)
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
