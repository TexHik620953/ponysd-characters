package application

import (
	"ponysd-characters/infra/controllers/characters"
	charactersimages "ponysd-characters/infra/controllers/charactersImages"
	"ponysd-characters/infra/controllers/glossary"
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

		charsPublic := charsGroup.Group("")
		// Public
		{
			charsPublic.GET("/:id", charCtrl.PublicGet)
			charsPublic.GET("/", charCtrl.PublicListCharacters)
		}

		charsPrivate := charsGroup.Group("")
		charsPrivate.Use(customMiddleware.NewUserIDAuth())
		// Private
		{
			charsPrivate.GET("/:id", charCtrl.Get)
			charsPrivate.GET("/", charCtrl.ListUserCharacters)
			charsPrivate.DELETE("/:id", charCtrl.Delete)
			charsPrivate.POST("/", charCtrl.Create)
		}
	}

	// Glossary
	{
		glosCtrl := glossary.New(app.services)
		glosGroup := api.Group("/glossary/:type")

		glosGroup.GET("/", glosCtrl.ListRecords)
		glosGroup.GET("/:local", glosCtrl.ListRecordsLocal)
	}

	// Characters image
	{
		charImgCtrl := charactersimages.New(app.services)
		charGroup := api.Group("/images/:char_id")

		charGroup.GET("/", charImgCtrl.ListCharImages)
		charGroup.GET("/:id", charImgCtrl.GetCharImage)
		charGroup.POST("/", charImgCtrl.Create)
	}

}
