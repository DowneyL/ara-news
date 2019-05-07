package validators

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	entrans "gopkg.in/go-playground/validator.v9/translations/en"
	jatrans "gopkg.in/go-playground/validator.v9/translations/ja"
	zhtrans "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
)

var (
	uni       *ut.UniversalTranslator
	Validator UniversalValidator
)

type UniversalValidator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func InitValidate() UniversalValidator {
	Validator.Validate = validator.New()

	// 验证器字段根据标签 form 或者 json 的名字定
	Validator.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		getFieldName := func(str string) string {
			return strings.SplitN(fld.Tag.Get(str), ",", 2)[0]
		}
		name := getFieldName("json")
		if name == "" {
			name = getFieldName("form")
		}
		if name == "-" {
			return ""
		}
		return name
	})

	return Validator
}

func InitUniversalValidator(lang string) UniversalValidator {
	switch lang {
	case "en-US":
		return registerEnTranslation()
	case "zh-CN":
		return registerZhTranslation()
	case "ja":
		return registerJaTranslation()
	default:
		return registerEnTranslation()
	}
}

func registerEnTranslation() UniversalValidator {
	enTr := en.New()
	uni = ut.New(enTr, enTr)
	trans, _ := uni.GetTranslator("en")
	Validator.Trans = trans
	_ = entrans.RegisterDefaultTranslations(Validator.Validate, Validator.Trans)
	return Validator
}

func registerZhTranslation() UniversalValidator {
	zhTr := zh.New()
	uni = ut.New(zhTr, zhTr)
	trans, _ := uni.GetTranslator("zh")
	Validator.Trans = trans
	_ = zhtrans.RegisterDefaultTranslations(Validator.Validate, Validator.Trans)
	return Validator
}

func registerJaTranslation() UniversalValidator {
	jaTr := ja.New()
	uni = ut.New(jaTr, jaTr)
	trans, _ := uni.GetTranslator("ja")
	Validator.Trans = trans
	_ = jatrans.RegisterDefaultTranslations(Validator.Validate, Validator.Trans)
	return Validator
}
