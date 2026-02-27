package middlewares

import (
	"ponysd-characters/pkg/httperr"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

const AuthHeader = "X-User-ID"

// Simple auth that checks whether X-USer-ID is present and is a valid userID
func NewUserIDAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// Get User-ID header
			userID := c.Request().Header.Get(AuthHeader)
			if len(userID) == 0 {
				return httperr.ErrUnauthorized("missing X-User-ID", nil)
			}

			id, err := uuid.Parse(userID)
			if err != nil {
				return httperr.ErrUnauthorized("failed to parse uuid from userID", err)
			}
			c.Set("User", id)
			return next(c)
		}
	}
}
