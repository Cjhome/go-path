package requests

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans ut.Translator
)

func init() {
	// 注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	// 获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

// Translate 翻译错误信息
func Translate(err error) map[string][]string {
	result := make(map[string][]string)
	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		fmt.Println("err", err)
		fmt.Println("err", err.Field())
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	fmt.Println("result", result)
	return result
}