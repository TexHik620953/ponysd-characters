package characters

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

func (ctrl *Controller) ListUserCharacters(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	_ = user
	return nil
}
func (ctrl *Controller) PublicListCharacters(c *echo.Context) error {
	return nil
}

func (ctrl *Controller) PublicGet(c *echo.Context) error {
	characterID := c.Param("id")

	_ = characterID
	return nil
}

func (ctrl *Controller) Get(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	characterID := c.Param("id")

	_ = characterID
	_ = user
	return nil
}

func (ctrl *Controller) Delete(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	characterID := c.Param("id")

	_ = characterID
	_ = user
	return nil
}

func (ctrl *Controller) Create(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	_ = user
	return nil
}
