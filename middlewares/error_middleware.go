package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"io"
	"unicode"
)

var (
	MsgInvalidJson     = "Invalid JSON format"
	MsgInvalidJsonType = func(e *json.UnmarshalTypeError) string {
		return "Expected " + e.Value + " but given type is " + e.Type.String() + " in JSON"
	}
	MsgValidationFailed      = "Invalid inputs"
	MsgValidationFieldFailed = func(e *validator.FieldError) string {
		// TODO i18n
		switch e.Tag {
		case "required":
			return "Input required"
		}
		return "Invalid input format [" + e.Tag + "]"
	}
)

type ClientFieldError struct {
	Code    string `json:"code,omitempty"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type ClientError struct {
	Message string             `json:"message,omitempty"`
	Errors  []ClientFieldError `json:"errors,omitempty"`
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.ByType(gin.ErrorTypeBind | gin.ErrorTypePublic).Last()

			if err != nil {
				clientError := &ClientError{}
				switch err.Err {
				case io.EOF:
					clientError.Message = MsgInvalidJson
				default:
					switch err.Err.(type) {
					case *json.SyntaxError:
						clientError.Message = MsgInvalidJson
					case *json.UnmarshalTypeError:
						clientError.Message = MsgInvalidJsonType(err.Err.(*json.UnmarshalTypeError))
					case validator.ValidationErrors:
						clientError.Message = MsgValidationFailed
						clientFieldErrors := []ClientFieldError{}
						for _, fieldErr := range err.Err.(validator.ValidationErrors) {
							clientFieldError := ClientFieldError{}
							clientFieldError.Code = fieldErr.Tag
							clientFieldError.Field = toSnakeCase(fieldErr.Field)
							clientFieldError.Message = MsgValidationFieldFailed(fieldErr)
							clientFieldErrors = append(clientFieldErrors, clientFieldError)
						}
						clientError.Errors = clientFieldErrors
					default:
						clientError.Message = err.Error()
					}
					c.JSON(-1, clientError) // -1 == not override the current error code
				}
			}
		}
	}
}

// https://gist.github.com/elwinar/14e1e897fdbe4d3432e1
func toSnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
