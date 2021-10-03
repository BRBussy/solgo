package validation

import (
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type CustomTranslation struct {
	Tag                string
	Translation        string
	Override           bool
	CustomRegisterFunc validator.RegisterTranslationsFunc
	CustomTransFunc    validator.TranslationFunc
}

func registerCustomTranslations(v *validator.Validate, trans ut.Translator, translations []CustomTranslation) (err error) {

	for _, t := range translations {
		if t.CustomTransFunc != nil && t.CustomRegisterFunc != nil {
			err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisterFunc, t.CustomTransFunc)
		} else if t.CustomTransFunc != nil && t.CustomRegisterFunc == nil {
			err = v.RegisterTranslation(t.Tag, trans, registrationFunc(t.Tag, t.Translation, t.Override), t.CustomTransFunc)
		} else if t.CustomTransFunc == nil && t.CustomRegisterFunc != nil {
			err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisterFunc, translateFunc)
		} else {
			err = v.RegisterTranslation(t.Tag, trans, registrationFunc(t.Tag, t.Translation, t.Override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registerDefaultOverrideTranslations(v *validator.Validate, trans ut.Translator) (err error) {
	translations := []CustomTranslation{
		{
			Tag:         "required_without",
			Translation: "{0} is required when {1} is not present",
			Override:    false,
			CustomTransFunc: func(ut ut.Translator, fe validator.FieldError) string {

				t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
				if err != nil {
					return fe.(error).Error()
				}

				return t
			},
		},
	}

	for _, t := range translations {
		if t.CustomTransFunc != nil && t.CustomRegisterFunc != nil {
			err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisterFunc, t.CustomTransFunc)
		} else if t.CustomTransFunc != nil && t.CustomRegisterFunc == nil {
			err = v.RegisterTranslation(t.Tag, trans, registrationFunc(t.Tag, t.Translation, t.Override), t.CustomTransFunc)
		} else if t.CustomTransFunc == nil && t.CustomRegisterFunc != nil {
			err = v.RegisterTranslation(t.Tag, trans, t.CustomRegisterFunc, translateFunc)
		} else {
			err = v.RegisterTranslation(t.Tag, trans, registrationFunc(t.Tag, t.Translation, t.Override), translateFunc)
		}

		if err != nil {
			return
		}
	}

	return
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return

	}

}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {

	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}

	return t
}
