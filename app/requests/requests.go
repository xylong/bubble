package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate 参数验证
func Validate(ctx *gin.Context, obj interface{}, handler ValidatorFunc) (gin.H, bool) {
	// 1.解析请求参数
	if err := ctx.ShouldBind(obj); err != nil {
		return gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式",
			"error":   err.Error(),
		}, false
	}

	// 2.验证参数
	if errs := handler(obj, ctx); len(errs) > 0 {
		return gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		}, false
	}

	return nil, true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
