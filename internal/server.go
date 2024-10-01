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

type PageData struct {
	Title     string
	UserName  string
	UserEmail string
	UserItems []UserItem
	RouteItem []Route

	DropdownTitle string
	DropdownItems []Route
}

// TODO: replace with auth service
type UserItem struct {
	ItemText string
	ItemURL  string
}

type Route struct {
	RouteTitle  string
	RouteURL    string
	RouteTarget string
}

type DropdownItem struct {
	ItemURL    string
	ItemTarget string
	ItemText   string
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
	e.Static("/static", "public/static")

	data := PageData{
		Title:     "Quiz App",
		UserName:  "Mazzi",
		UserEmail: "andre.mazzq@gamil.com",
		UserItems: []UserItem{
			{ItemText: "info", ItemURL: "/auth/info"},
			{ItemText: "sign out", ItemURL: "/auth/sign-out"},
		},
		RouteItem: []Route{
			{RouteTitle: "Questions", RouteTarget: "left-card", RouteURL: "/form/findAll"},
			{RouteTitle: "Quizs", RouteTarget: "left-card", RouteURL: "/quiz/findAll"},
		},
		DropdownTitle: "Create Question",
		DropdownItems: []Route{
			{RouteTitle: "Open Question", RouteTarget: "right-card", RouteURL: "/form/open-question/create-page"},
			{RouteTitle: "Closed Question", RouteTarget: "right-card", RouteURL: "/form/closed-question/create-page"},
		},
	}

	e.GET("/*", func(c echo.Context) error { return c.Render(http.StatusNotFound, "error.html", "404 not found") })
	e.GET("/", func(c echo.Context) error { return c.Render(http.StatusOK, "home.html", data) })
	adminRoute := e.Group("/auth")
	formRoute := e.Group("/form")
	admin.Route(adminRoute)
	form.Route(formRoute)

	common.Logger.LogInfo().Msg(e.Start(":" + os.Getenv("APP_PORT")).Error())
}
