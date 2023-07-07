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
	})
}

func (c *SignupController) IsPhoneExist(ctx *gin.Context) interface{} {
	// 解析请求参数
	request := &requests.SignupPhoneExistRequest{}
	if err := ctx.ShouldBind(request); err != nil {
		return err.Error()
	}

	// 表单验证
	if errs := requests.ValidateSignupPhoneExist(request, ctx); len(errs) > 0 {
		return errs
	}

	return user.IsPhoneExist(request.Phone)
}
