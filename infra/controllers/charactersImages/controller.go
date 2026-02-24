package charactersimages

import (
	"ponysd-characters/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type Controller struct {
	services services.ServicesContext
}

func New(svc services.ServicesContext) *Controller {
	return &Controller{services: svc}
}

// List images indexes of character
func (ctrl *Controller) ListCharImages(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	charID := c.Param("char_id")

	_ = user
	_ = charID
	return nil
}

// Get exact image binary
func (ctrl *Controller) GetCharImage(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	charID := c.Param("char_id")
	imageID := c.Param("id")

	_ = user
	_ = charID
	_ = imageID
	return nil
}
