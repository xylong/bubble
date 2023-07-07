package auth

import (
	v1 "bubble/app/http/controllers/api/v1"
	"bubble/app/models/user"
	"bubble/app/requests"
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
		auth.POST("/signup/email/exist", c.IsEmailExist)
	})
}

func (c *SignupController) IsPhoneExist(ctx *gin.Context) interface{} {
	// 解析请求参数
	request := &requests.SignupPhoneExistRequest{}
	if result, ok := requests.Validate(ctx, request, requests.ValidateSignupPhoneExist); !ok {
		return result
	}

	return user.IsPhoneExist(request.Phone)
}

func (c *SignupController) IsEmailExist(ctx *gin.Context) interface{} {
	request := &requests.SignupEmailExistRequest{}
	if result, ok := requests.Validate(ctx, request, requests.ValidateSignupEmailExist); !ok {
		return result
	}

	return user.IsEmailExist(request.Email)
}
