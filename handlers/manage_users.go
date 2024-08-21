package handlers

import (
	"net/http"
	"strconv"

	"example.com/goserver/models"
	"example.com/goserver/views"
)

func ExampleUsers() []models.User {
	return []models.User{
		{Username: "Bob", Email: "bob@example.com"},
		{Username: "George", Email: "george@example.com"},
		{Username: "Mary", Email: "mary@example.com"},
		{Username: "Jane", Email: "jane@example.com"},
		{Username: "Andy", Email: "andy@example.com"},
	}
}

func ManageUsersHandler(w http.ResponseWriter, r *http.Request) {
	hxReq, _ := strconv.ParseBool(r.Header.Get("HX-Request"))
	users := ExampleUsers()
	if !hxReq {
		views.Base("Users", views.UsersTable(users)).Render(r.Context(), w)
	} else {
		views.UsersTable(users).Render(r.Context(), w)
	}
}
