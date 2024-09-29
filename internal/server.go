package internal

import (
	"form_management/common"
	"form_management/internal/admin"
	"form_management/internal/form"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Server() {
	e := echo.New()

	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/**/*.html")),
	}
	e.Renderer = renderer
	e.Static("/", "public")

	e.GET("/*", func(c echo.Context) error { return c.Render(http.StatusNotFound, "error.html", "404 not found") })
	e.GET("/", func(c echo.Context) error { return c.Render(http.StatusOK, "home.html", "Forms App") })
	adminRoute := e.Group("/admin")
	formRoute := e.Group("/form")
	admin.Route(adminRoute)
	form.Route(formRoute)

	common.Logger.LogInfo().Msg(e.Start(":" + os.Getenv("APP_PORT")).Error())
}
