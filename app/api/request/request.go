package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type requester struct{}

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

func init() {
	setZh()
}

func setZh() {
	Uni = ut.New(zh.New())
	Validate = validator.New()
	Trans, _ = Uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(Validate, Trans)
}

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
