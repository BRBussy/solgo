package validation

import (
	"fmt"
	enLocales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/rs/zerolog/log"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/en"
	"sort"
	"strings"
)

// GoPlaygroundValidator is a go-playground/validator.v9
// implementation of the Validator interface.
type GoPlaygroundValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

// Validate is the default implementation of the ValidateRequest service. The go-playground validator and
// universal translator is utilised. In future the implementation should allow for the language to be selected.
func (v *GoPlaygroundValidator) Validate(request interface{}) error {
	var reasons []string
	// check for validation errors and append to reasons
	if err := v.validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := validationErrors.Translate(v.translator)
		for key, value := range errorMessages {
			reasons = append(reasons, fmt.Sprintf("'%s' %s", key, value))
		}
	}
	// check reasons
	if len(reasons) > 0 {
		sort.Slice(reasons, func(i, j int) bool {
			return reasons[i] > reasons[j]
		})
		return ErrValidationFailed{Reason: strings.Join(reasons, ";")}
	}

	return nil
}

func NewValidator(
	customValidations []CustomValidation,
	customTranslations []CustomTranslation,
) (*GoPlaygroundValidator, error) {
	e := enLocales.New()
	uni := ut.New(e, e)
	trans, _ := uni.GetTranslator("en")
	vv := validator.New()

	// register default translations
	if err := en.RegisterDefaultTranslations(vv, trans); err != nil {
		log.Error().Err(err).Msg("could not register default translations")
		return nil, fmt.Errorf(
			"could not register default translations: %w",
			err,
		)
	}

	// register default translation overrides
	if err := registerDefaultOverrideTranslations(vv, trans); err != nil {
		log.Error().Err(err).Msg("could not register default override translations")
		return nil, fmt.Errorf(
			"could not register default override translations: %w",
			err,
		)
	}

	// register custom translations
	if err := registerCustomTranslations(vv, trans, customTranslations); err != nil {
		log.Error().Err(err).Msg("could not register custom translations")
		return nil, fmt.Errorf(
			"could not register custom translations: %w",
			err,
		)
	}

	// register custom validation functions
	if err := registerCustomValidations(vv, customValidations); err != nil {
		log.Error().Err(err).Msg("could not register custom validations translations")
		return nil, fmt.Errorf(
			"could not register custom validations translations: %w",
			err,
		)
	}

	return &GoPlaygroundValidator{
		validate:   vv,
		translator: trans,
	}, nil
}

func MustNewGoPlaygroundValidator(
	customValidations []CustomValidation,
	customTranslations []CustomTranslation,
) *GoPlaygroundValidator {
	v, err := NewValidator(customValidations, customTranslations)
	if err != nil {
		log.Fatal().Err(err).Msg("could not create go playground validator")
	}

	return v
}
