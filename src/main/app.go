package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"main/api"
)

func init()  {
	e := CreateEcho()
	http.Handle("/", e)
}

func CreateEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")

	api.InitNotes(*v1, "/notes")

	return e
}