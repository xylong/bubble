package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
)

type UserCtrl struct {
}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

func (c *UserCtrl) Name() string {
	return "UserCtrl"
}

func (c *UserCtrl) Route(group *bingo.Group) {
	group.POST("register", c.register)
}

func (c *UserCtrl) register(ctx *gin.Context) string {
	return "注册"
}
