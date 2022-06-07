package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Printf("Hit the home page endpoint")
}

// userPage
func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("hit the users page endpoint")
	json.NewEncoder(w).Encode(users)
	fmt.Fprintf(w, "Welcome to the home page")
}

func main() {
	// connectDb()
	connectMYSQL()
}

// func connectDb() {
// 	e := echo.New()

// 	d := db.DBConnect()
// 	h := users.NewHandler(user.NewUserModel(d))

// 	e.GET("/users", h.GetIndex)
// 	e.GET("/users/:id", h.GetDetail)

// 	e.Logger.Fatal(e.Start(":2250"))
// }

func connectMYSQL() {
	fmt.Println("Welcome conenctmyFun")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	http.ListenAndServe(":1150", nil)
}
