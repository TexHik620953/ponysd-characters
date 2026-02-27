package glossary

import (
	"net/http"
	"ponysd-characters/services"

	"github.com/labstack/echo/v5"
)

type Controller struct {
	services services.ServicesContext
}

func New(svc services.ServicesContext) *Controller {
	return &Controller{services: svc}
}

func (ctrl *Controller) ListTypes(c *echo.Context) error {
	svc := ctrl.services.GlossaryService()

	records, err := svc.ListTypes(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, records)
}

func (ctrl *Controller) ListRecords(c *echo.Context) error {
	typ := c.Param("type")

	svc := ctrl.services.GlossaryService()

	records, err := svc.ListRecordsByType(c.Request().Context(), typ)
	if err != nil {
		return err
	}

	result := make([]*GlossaryRecordResponse, len(records))
	for i, r := range records {
		result[i] = mapGlossaryRecordToResponse(&r)
	}
	return c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) ListRecordsLocal(c *echo.Context) error {
	typ := c.Param("type")
	local := c.Param("local")

	svc := ctrl.services.GlossaryService()

	records, err := svc.ListRecordLocal(c.Request().Context(), typ, local)
	if err != nil {
		return err
	}

	result := make([]*GlossaryRecordResponse, len(records))
	for i, r := range records {
		result[i] = mapGlossaryRecordLocalToResponse(&r)
	}
	return c.JSON(http.StatusOK, result)
}
