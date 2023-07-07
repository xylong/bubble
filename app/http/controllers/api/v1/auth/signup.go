package auth

import (
	v1 "bubble/app/http/controllers/api/v1"
	"bubble/app/models/user"
	"github.com/gin-gonic/gin"
	"github.com/xylong/bingo"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseApiController
}

func NewSignupController() *SignupController {
	return &SignupController{}
}

func (c *SignupController) Name() string {
	return "SignupController"
}

func (c *SignupController) Route(group *bingo.Group) {
	group.Group("auth", func(auth *bingo.Group) {
		auth.POST("/signup/phone/exist", c.IsPhoneExist)
	})
}

func (c *SignupController) IsPhoneExist(ctx *gin.Context) interface{} {
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	req := &PhoneExistRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		return err.Error()
	}

	return user.IsPhoneExist(req.Phone)
}
