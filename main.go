package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/", getQuestions)
	e.Logger.Fatal(e.Start(":8080"))
}

func getQuestions(c echo.Context) error {
	return c.String(http.StatusOK, "Sas here some questions")

}
