package util

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func TranslateErroValidation(validate *validator.Validate, err error) error {
	if err == nil {
		return nil
	}
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	errsTrans := err.(validator.ValidationErrors).Translate(trans)

	for _, errTrans := range errsTrans {
		return errors.New(errTrans)
	}
	return nil
}
