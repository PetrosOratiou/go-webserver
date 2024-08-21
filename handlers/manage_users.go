package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"example.com/goserver/models"
	"example.com/goserver/views"
	_ "github.com/mattn/go-sqlite3"
)

func ExampleUsers() []models.User {
	return []models.User{
		{Id: 1, Username: "Bob", Email: "bob@example.com"},
		{Id: 2, Username: "George", Email: "george@example.com"},
		{Id: 3, Username: "Mary", Email: "mary@example.com"},
		{Id: 4, Username: "Jane", Email: "jane@example.com"},
		{Id: 5, Username: "Andy", Email: "andy@example.com"},
	}
}

type Users struct {
	List []models.User
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUsersHandler")
	hxReq, _ := strconv.ParseBool(r.Header.Get("HX-Request"))
	users, _ := SelectUsers()
	if !hxReq {
		views.Base("Users", views.UsersTable(users)).Render(r.Context(), w)
	} else {
		views.UsersTable(users).Render(r.Context(), w)
	}
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.ParseForm() != nil {
		w.WriteHeader(400)
		return
	}
	u := &models.User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
	}
	u, err := InsertUser(u)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	views.UserTableRow(u).Render(r.Context(), w)
}

func InsertUser(u *models.User) (*models.User, error) {
	db, err := sql.Open("sqlite3", "sql/local.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	row, err := tx.Exec(`INSERT INTO users (username, email) VALUES (?, ?)`, u.Username, u.Email)
	if err != nil {
		return nil, err
	}
	userid, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	u.Id = int(userid)
	return u, nil
}

func SelectUsers() ([]*models.User, error) {
	var db *sql.DB
	var err error
	if db, err = sql.Open("sqlite3", "sql/local.db"); err != nil {
		return nil, err
	}
	var rows *sql.Rows
	if rows, err = db.Query(`SELECT rowid, username, email FROM users`); err != nil {
		return nil, err
	}
	users := []*models.User{}
	for rows.Next() {
		u := models.User{}
		if err = rows.Scan(&u.Id, &u.Username, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
