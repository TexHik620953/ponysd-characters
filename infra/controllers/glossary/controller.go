package glossary

import (
	"ponysd-characters/services"

	"github.com/labstack/echo/v5"
)

type Controller struct {
	services services.ServicesContext
}

func New(svc services.ServicesContext) *Controller {
	return &Controller{services: svc}
}

func (ctrl *Controller) ListRecords(c *echo.Context) error {
	typ := c.Param("type")

	_ = typ
	return nil
}

func (ctrl *Controller) ListRecordsLocal(c *echo.Context) error {
	typ := c.Param("type")
	local := c.Param("local")

	_ = typ
	_ = local
	return nil
}
