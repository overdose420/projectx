package users

import (
	"database/sql"
	"net/http"

	user "overp/models"

	"github.com/labstack/echo"
)

type resultLists struct {
	Users []user.User `json:"users"`
}

type handler struct {
	UserModel user.UserModelImpl
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewHandler(u user.UserModelImpl) *handler {
	return &handler{u}
}

func (h *handler) GetIndex(c echo.Context) error {
	lists := h.UserModel.FindAll()
	u := &resultLists{
		Users: lists,
	}
	return c.JSON(http.StatusOK, u)
}

func (h *handler) GetDetail(c echo.Context) error {
	id := c.Param("id")
	u := h.UserModel.FindByID(id)
	return c.JSON(http.StatusOK, u)
}

func getUsers() []*User {
	// Open up db connection
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// check if error with database connection
	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []*User
	for results.Next() {
		var u User

		// scan each result
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}

	return users
}
