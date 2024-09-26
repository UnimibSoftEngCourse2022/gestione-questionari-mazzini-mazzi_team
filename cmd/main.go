package main

import (
	"form_management/common"
	"form_management/internal/admin"
	"form_management/internal/form"
	"html/template"
	"io"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	err := godotenv.Load()
	if err != nil {
		common.Logger.Fatal().Msg("||=> Error loading .env file")
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/**/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "static")

	adminRoute := e.Group("/admin")
	formRoute := e.Group("/form")
	admin.Route(adminRoute)
	form.Route(formRoute)

	common.Logger.LogInfo().Msg(e.Start(":" + os.Getenv("APP_PORT")).Error())
}
