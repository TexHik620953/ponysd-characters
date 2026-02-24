package characters

import (
	"encoding/json"
	"net/http"
	"ponysd-characters/pkg/httperr"
	"ponysd-characters/services"
	"ponysd-characters/services/character"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

const PageSize = int64(50)

type Controller struct {
	services services.ServicesContext
}

func New(svc services.ServicesContext) *Controller {
	return &Controller{services: svc}
}

// List all characters of user
func (ctrl *Controller) ListUserCharacters(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)

	svc := ctrl.services.CharacterService()
	characters, err := svc.ListUserCharacters(c.Request().Context(), user)
	if err != nil {
		return err
	}

	result := make([]*CharacterResponse, len(characters))
	for i, c := range characters {
		result[i] = mapCharacterToResponse(&c)
	}

	return c.JSON(http.StatusOK, result)
}

// List public characters
func (ctrl *Controller) PublicListCharacters(c *echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 32)
	if err != nil {
		return httperr.ErrBadRequest("page is invalid", err)
	}

	svc := ctrl.services.CharacterService()
	characters, err := svc.ListPublicCharacters(c.Request().Context(), int(PageSize), int(page*PageSize))

	result := make([]*CharacterResponse, len(characters))
	for i, c := range characters {
		result[i] = mapCharacterToResponse(&c)
	}

	return c.JSON(http.StatusOK, result)
}

func (ctrl *Controller) PublicGet(c *echo.Context) error {
	characterID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return httperr.ErrBadRequest("id is invalid", err)
	}

	svc := ctrl.services.CharacterService()

	char, err := svc.GetCharacter(c.Request().Context(), uuid.Nil, characterID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, mapCharacterToResponse(char))
}

func (ctrl *Controller) Get(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	characterID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return httperr.ErrBadRequest("id is invalid", err)
	}

	svc := ctrl.services.CharacterService()

	char, err := svc.GetCharacter(c.Request().Context(), user, characterID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, mapCharacterToResponse(char))
	return nil
}

func (ctrl *Controller) Delete(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)
	characterID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return httperr.ErrBadRequest("id is invalid", err)
	}

	svc := ctrl.services.CharacterService()
	svc.DeleteCharacter(c.Request().Context(), user, characterID)

	return c.NoContent(http.StatusNoContent)
}

func (ctrl *Controller) Create(c *echo.Context) error {
	user := c.Get("User").(uuid.UUID)

	var req CreateCharacterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return httperr.ErrBadRequest("invalid body", err)
	}

	svc := ctrl.services.CharacterService()

	char := &character.Character{
		OwnerID:     user,
		Name:        req.Name,
		Biography:   req.Biography,
		Nationality: req.Nationality,
		Age:         req.Age,
		Body:        req.Body,
		Breast:      req.Breast,
		Butt:        req.Butt,
		EyesColor:   req.EyesColor,
		HairStyle:   req.HairStyle,
		HairColor:   req.HairColor,
	}

	svc.CreateCharacter(c.Request().Context(), char, int64(req.Seed))

	return c.JSON(http.StatusCreated, char)
}
