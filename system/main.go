package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
)

type User struct {
	Name string
	Password string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "goshop"
	password = "1234"
)

var (
	db  *sql.DB
	err error
)

// func home_page(w http.ResponseWriter, r *http.Request) {
// 	// bob := User{"bob", "asd", "1234"}
// 	temp, _ := template.ParseFiles("templates/home_page.html")
// 	temp.Execute(w, nil)
// }

func registerHandler(w http.ResponseWriter, r *http.Request) {

	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	CheckError(err)
	defer db.Close()

	if r.Method != http.MethodPost {
		// Render registration form
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Process registration form submission
	username := r.FormValue("name")
	password := r.FormValue("password")

	// Connect to the database
	db, err = sql.Open("postgres", db_connection)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if the username already exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name=?", username).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count > 0 {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, "Username already exists")
		return
	}

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Registration successful
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/register", registerHandler)
	// http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8181", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	CheckError(err)
	defer db.Close()

	if r.Method != http.MethodPost {
		// Render login form
		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Process login form submission
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Connect to the database
	db, err = sql.Open("postgres", db_connection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check if the username and password are valid
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username=? AND password=?", username, password).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w, "Invalid username or password")
		return
	}

	// Login successful
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// func contacts_page(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Contacts page")
// }

// func handleRequest() {
// 	http.HandleFunc("/", home_page)
// 	// http.HandleFunc("/contacts/", contacts_page)
// 	http.ListenAndServe(":4334", nil)
// }
