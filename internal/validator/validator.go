package validator

import (
	"net/http"
	"strings"

	e "auth/internal/transport/http/error"

	"github.com/go-playground/validator"
)

type ValidateData struct {
	e.HTTPError
	Fields []string `json:"fields"`
}

// Функция обработки ошибок валидации
func HandleValidationErrors(w http.ResponseWriter, err error) (*ValidateData, bool) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var vd ValidateData

		vd.Fields = []string{}

		for _, fieldErr := range validationErrors {
			vd.Fields = append(vd.Fields, strings.ToLower(fieldErr.Field()))
		}
		vd.Err = "Не все поля заполнены"
		return &vd, true
	}

	return nil, false
}
