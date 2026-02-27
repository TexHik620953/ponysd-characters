package tasks

import (
	"bytes"
	"image/png"
	"net/http"
	"ponysd-characters/pkg/httperr"
	"ponysd-characters/services"
	"ponysd-characters/services/task"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type Controller struct {
	services services.ServicesContext
}

func New(servicesCtx services.ServicesContext) *Controller {
	return &Controller{services: servicesCtx}
}

func (ctrl *Controller) PullTask(c *echo.Context) error {
	task, err := ctrl.services.TasksService().PullPendingTask(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &Task{
		ID:             task.ID,
		PositivePrompt: task.PositivePrompt,
		NegativePrompt: task.NegativePrompt,
		Seed:           task.Seed,
	})
}

func (ctrl *Controller) PushResult(c *echo.Context) error {
	taskID, err := uuid.Parse(c.Param("task_id"))
	if err != nil {
		return httperr.ErrBadRequest("id is invalid", err)
	}

	img, err := png.Decode(c.Request().Body)
	if err != nil {
		return httperr.ErrBadRequest("body image is invalid", err)
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return httperr.ErrInternalServerError("failed to encode image", err)
	}

	fileHash, err := ctrl.services.FileStorageService().StoreFile(c.Request().Context(), buf.Bytes())
	if err != nil {
		return err
	}

	tx, services, err := ctrl.services.BeginTx(c.Request().Context())
	if err != nil {
		return err
	}
	defer tx.Rollback(c.Request().Context())

	taskInfo, err := services.TasksService().GetTaskByID(c.Request().Context(), taskID)
	if err != nil {
		return err
	}

	err = services.TasksService().SetTaskStatus(c.Request().Context(), taskID, task.StatusDone)
	if err != nil {
		return err
	}

	_, err = services.ImageService().InsertImage(c.Request().Context(), taskInfo.UserID, taskInfo.CharacterID, fileHash)
	if err != nil {
		return err
	}

	return tx.Commit(c.Request().Context())
}
