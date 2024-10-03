package internal

import (
	common "form_management/common/logger"
	"form_management/internal/auth"
	"form_management/internal/form"
	"form_management/internal/quiz"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type AuthPage struct {
	Title   string
	HtmxURL string
}

type PageData struct {
	Title     string
	UserName  string
	UserEmail string
	UserItems []UserItem
	RouteItem []Route

	DropdownTitle string
	DropdownItems []Route

	DropdownInputItem       string
	DropdownInputTitle      string
	DropdownInputItemTarget string
	DropdownInputItemURL    string
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

var data = PageData{
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
	DropdownInputTitle:      "Create New Quiz",
	DropdownInputItem:       "Insert Name of New Quiz",
	DropdownInputItemURL:    "/quiz/create",
	DropdownInputItemTarget: "quizs",
}

var RegisterData = AuthPage{
	Title:   "Form App",
	HtmxURL: "/auth/register/user",
}

var LoginData = AuthPage{
	Title:   "Form App",
	HtmxURL: "/auth/login/user",
}

func Server() {
	e := echo.New()

	common.NewLogger()
	e.Use(common.LoggingMiddleware)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/**/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "public/static")

	e.GET("/*", func(c echo.Context) error { return c.Render(http.StatusNotFound, "error.html", "404 not found") })

	// TODO: create specific handler
	// auth page !! [ anche loro nell'handler di home ??]
	e.GET("/login", func(c echo.Context) error { return c.Render(http.StatusOK, "login.page.html", LoginData) })
	e.GET("/register", func(c echo.Context) error { return c.Render(http.StatusOK, "register.page.html", RegisterData) })
	e.GET("/", func(c echo.Context) error { return c.Render(http.StatusOK, "home.page.html", data) }, auth.AuthMiddleware)

	authRoute := e.Group("/auth")
	formRoute := e.Group("/form", auth.AuthMiddleware)
	quizRoute := e.Group("/quiz", auth.AuthMiddleware)
	auth.Route(authRoute)
	form.Route(formRoute)
	quiz.Route(quizRoute)

	common.Logger.LogInfo().Msg(e.Start(":" + os.Getenv("APP_PORT")).Error())
}
