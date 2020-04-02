package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
	"log"
)

type requester struct{}

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	setZhInit()
	registerInit()
}

func setZhInit() {
	Uni = ut.New(zh.New())
	Validate = validator.New()
	Trans, _ = Uni.GetTranslator("zh")
	if err := zhTranslations.RegisterDefaultTranslations(Validate, Trans); err != nil {
		log.Println(err)
	}
}

//基础的中文验证
func (r *requester) ValidateZH(c *gin.Context, s interface{}) error {
	if err := c.Bind(s); err != nil {
		return err
	}

	if err := Validate.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		return errors.New(errs[0].Translate(Trans))
	}

	return nil
}

func registerInit() {
	Validate.RegisterValidation("organizeMarketTypeSymbolUnique", organizeMarketTypeSymbolUnique)
	Validate.RegisterTranslation("organizeMarketTypeSymbolUnique", Trans, func(ut ut.Translator) error {
		return ut.Add("organizeMarketTypeSymbolUnique", "该交易所币对已添加", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("organizeMarketTypeSymbolUnique", fe.Field())
		return t
	})
}
