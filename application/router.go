package application

import (
	"ponysd-characters/infra/controllers/characters"
	charactersimages "ponysd-characters/infra/controllers/charactersImages"
	"ponysd-characters/infra/controllers/glossary"
	"ponysd-characters/infra/controllers/tasks"
	customMiddleware "ponysd-characters/infra/middlewares"

	"github.com/labstack/echo/v5/middleware"
)

func (app *Application) configureRoutes() {

	api := app.srv.Group("/api")
	api.Use(middleware.RequestLogger())
	api.Use(customMiddleware.ErrorMiddleware())

	// Characters CRUD
	{
		charCtrl := characters.New(app.services)

		charsGroup := api.Group("/characters")

		charsPrivate := charsGroup.Group("")
		charsPrivate.Use(customMiddleware.NewUserIDAuth())
		// Private
		{
			charsPrivate.GET("/:id", charCtrl.Get)
			charsPrivate.GET("", charCtrl.ListUserCharacters)
			charsPrivate.DELETE("/:id", charCtrl.Delete)
			charsPrivate.POST("", charCtrl.Create)
		}

		charsPublic := charsGroup.Group("/public")
		// Public
		{
			charsPublic.GET("/:id", charCtrl.PublicGet)
			charsPublic.GET("", charCtrl.PublicListCharacters)
		}
	}

	// Glossary
	{
		glosCtrl := glossary.New(app.services)
		glosGroup := api.Group("/glossary")

		glosGroup.GET("", glosCtrl.ListTypes)
		glosGroup.GET("/:type", glosCtrl.ListRecords)
		glosGroup.GET("/:type/:local", glosCtrl.ListRecordsLocal)
	}

	// Characters image
	{
		charImgCtrl := charactersimages.New(app.services)

		charGroup := api.Group("/images/:char_id")
		charGroup.Use(customMiddleware.NewUserIDAuth())

		charGroup.GET("/:id", charImgCtrl.GetCharImage)
		charGroup.GET("", charImgCtrl.ListCharImages)
		charGroup.POST("", charImgCtrl.Create)
	}

	// Tasks
	{
		tasksCtrl := tasks.New(app.services)

		tasksGroup := api.Group("/tasks")

		tasksGroup.GET("", tasksCtrl.PullTask)
		tasksGroup.POST("/:task_id", tasksCtrl.PushResult)
	}

}
