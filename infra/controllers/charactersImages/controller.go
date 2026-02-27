package charactersimages

import (
	"encoding/json"
	"net/http"
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
	charID, err := uuid.Parse(c.Param("char_id"))
	if err != nil {
		return err
	}

	svc := ctrl.services.ImageService()
	images, err := svc.ListCharacterImages(c.Request().Context(), user, charID)
	if err != nil {
		return err
	}

	result := make([]*FileInfoResponse, len(images))
	for i, img := range images {
		result[i] = mapServiceFileInfoToResponse(&img)
	}

	return c.JSON(http.StatusOK, result)
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

// Enqueue image generation
func (ctrl *Controller) Create(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	charID := c.Param("char_id")

	var req GenerateImageRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return err
	}

	_ = user
	_ = charID

	// TODO: Push task here to queue, and return task info and id
	return nil
}

// Ask how many tokens will this generation cost
func (ctrl *Controller) ProbeCreate(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	charID := c.Param("char_id")

	var req GenerateImageRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return err
	}

	_ = user
	_ = charID

	return c.JSON(http.StatusOK, &ProbeResponse{
		TokensRequired: 10,
	})
}
