package user

import (
	"log"

	. "framework.learning/handler"
	"framework.learning/model"
	"framework.learning/pkg/errno"
	"framework.learning/util"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	log.Printf("creating user: reqId: " + util.GetReqID(c))
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Printf("username is [%s], password is [%s]", r.Username, r.Password)

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
	}

	if err := u.Create(); err != nil {
		log.Print(err)
		SendResponse(c, errno.ErrDatabase, nil)
	}

	SendResponse(c, nil, CreateResponse{
		Username: r.Username,
	})
}
