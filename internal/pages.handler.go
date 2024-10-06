package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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

	QuizId string
}

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

func LoginPageHandler(c echo.Context) error {
	var LoginData = AuthPage{
		Title:   "Quiz App",
		HtmxURL: "/auth/login/user",
	}
	return c.Render(http.StatusOK, "login.page.html", LoginData)
}

func RegisternPageHandler(c echo.Context) error {
	var RegisterData = AuthPage{
		Title:   "Quiz App",
		HtmxURL: "/auth/register/user",
	}
	return c.Render(http.StatusOK, "register.page.html", RegisterData)
}

func HomePageHandler(c echo.Context) error {
	var data = PageData{
		Title:     "Quiz App",
		UserName:  "Mazzi",
		UserEmail: "andre.mazzq@gamil.com",
		UserItems: []UserItem{
			{ItemText: "logout", ItemURL: "/auth/logout"},
		},
		RouteItem: []Route{
			{RouteTitle: "Questions", RouteTarget: "body", RouteURL: "/questions"},
			{RouteTitle: "Quizs", RouteTarget: "body", RouteURL: "/quizs"},
		},
	}

	return c.Render(http.StatusOK, "home.page.html", data)
}

func QuestionPageHandler(c echo.Context) error {
	var data interface{} = PageData{
		Title:     "Quiz App",
		UserName:  "Utente",
		UserEmail: "",
		UserItems: []UserItem{
			{ItemText: "logout", ItemURL: "/auth/logout"},
		},
		RouteItem: []Route{
			{RouteTitle: "Questions", RouteTarget: "body", RouteURL: "/questions"},
			{RouteTitle: "Quizs", RouteTarget: "body", RouteURL: "/quizs"},
		},
		DropdownTitle: "Create Question",
		DropdownItems: []Route{
			{RouteTitle: "Open Question", RouteTarget: "right-card", RouteURL: "/form/open-question/create-page"},
			{RouteTitle: "Closed Question", RouteTarget: "right-card", RouteURL: "/form/closed-question/create-page"},
		},
	}
	return c.Render(http.StatusOK, "questions.page.html", data)
}

func QuizPageHandler(c echo.Context) error {
	var data interface{} = PageData{
		Title:     "Quiz App",
		UserName:  "Utente",
		UserEmail: "",
		UserItems: []UserItem{
			{ItemText: "logout", ItemURL: "/auth/logout"},
		},
		RouteItem: []Route{
			{RouteTitle: "Questions", RouteTarget: "body", RouteURL: "/questions"},
			{RouteTitle: "Quizs", RouteTarget: "body", RouteURL: "/quizs"},
		},
		DropdownInputTitle:      "Create New Quiz",
		DropdownInputItem:       "Insert Name of New Quiz",
		DropdownInputItemURL:    "/quiz/create",
		DropdownInputItemTarget: "quizs",
	}
	return c.Render(http.StatusOK, "quiz.page.html", data)
}

func QuizEditPageHandler(c echo.Context) error {

	quizID := c.QueryParam("quizID")
	var data interface{} = PageData{
		Title:     "Quiz App",
		UserName:  "Utente",
		UserEmail: "",
		UserItems: []UserItem{
			{ItemText: "logout", ItemURL: "/auth/logout"},
		},
		RouteItem: []Route{
			{RouteTitle: "Questions", RouteTarget: "body", RouteURL: "/questions"},
			{RouteTitle: "Quizs", RouteTarget: "body", RouteURL: "/quizs"},
		},
		QuizId: quizID,
	}

	return c.Render(http.StatusOK, "quiz-edit.page.html", data)
}
