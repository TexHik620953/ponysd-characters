package middlewares

import (
	"net/http"

	httperr "ponysd-characters/pkg/httperr"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

// ErrorMiddleware is a global middleware that:
//   - translates validation errors into 400 responses
//   - delegates httperr.HttpError instances to centralized JSON responses
//   - wraps unknown errors as 500 responses
func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if err := next(c); err != nil {
				// Validator errors -> 400 Bad Request
				if verrs, ok := err.(validator.ValidationErrors); ok {
					httpErr := httperr.BadRequest("validation error", verrs)
					status, body := httperr.AsHTTP(httpErr)
					_ = c.JSON(status, body)
					return nil
				}

				status, body := httperr.AsHTTP(err)
				if status == 0 {
					status = http.StatusInternalServerError
				}
				_ = c.JSON(status, body)
				return nil
			}
			return nil
		}
	}
}
