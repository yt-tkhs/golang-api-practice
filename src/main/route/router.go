package route

import (
	"github.com/labstack/echo"
	"main/api"
	"github.com/labstack/echo/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Gzip())

	v1 := e.Group("/api/v1")
	notes := v1.Group("/notes")
	{
		notes.GET("", api.GetAllNote())
		notes.GET("/:id", api.GetNote)
		notes.POST("", api.CreateNote)
	}

	return e
}