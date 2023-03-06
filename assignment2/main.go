package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "goshop"
	ps     = "1234"
)

var (
	db  *sql.DB
	err error
)

type User struct {
	Name     string
	Password string
}

type Item struct {
	ID       int
	Name     string
	Price    float32
	Raiting  float32
	Quantity int
}

func main() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/actions", actions)
	http.HandleFunc("/showall", showall)
	http.HandleFunc("/filter", filterItems)
	http.HandleFunc("/search", searchByName)
	// http.HandleFunc("/rating", giveRating)

	http.ListenAndServe(":8181", nil)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("templates/home_page.html")
	temp.Execute(w, nil)
}

func actions(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("templates/actions.html")
	temp.Execute(w, nil)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()
	if r.Method == "POST" {

		username := r.FormValue("username")
		pass := r.FormValue("password")

		_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2)`, username, pass)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "User %s created successfully!", username)
	} else {

		http.ServeFile(w, r, "templates/register.html")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name=$1 and password=$2", username, password).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count == 0 {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, "Invalid username or password")
		return
	}

	http.Redirect(w, r, "/actions", http.StatusSeeOther)
}

func showall(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT name, price, rating FROM store `)
	CheckError(err)

	defer rows.Close()
	var data []Item

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.Name, &i.Price, &i.Raiting)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}

	fmt.Fprintf(w, "<table>")
	fmt.Fprintf(w, " <h2> All items in store </h2><tr><th>Name</th><th>Price</th><th>Raiting</th></tr>")
	for _, i := range data {
		fmt.Fprintf(w, "<tr><td>%s</td><td>%v</td><td>%v</td></tr>", i.Name, i.Price, i.Raiting)
	}
	fmt.Fprintf(w, "</table>")
}

func filterItems(w http.ResponseWriter, r *http.Request) {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	rows, err := db.Query(`SELECT name, price, rating FROM store order by price`)
	CheckError(err)

	defer rows.Close()
	var data []Item

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.Name, &i.Price, &i.Raiting)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, i)
	}

	fmt.Fprintf(w, "<table>")
	fmt.Fprintf(w, " <h2> Filtering by price </h2> <tr><th>Name</th><th>Price</th><th>Raiting</th></tr>")
	for _, i := range data {
		fmt.Fprintf(w, "<tr><td>%s</td><td>%v</td><td>%v</td></tr>", i.Name, i.Price, i.Raiting)
	}
	fmt.Fprintf(w, "</table>")
}

func searchByName(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.New("search").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<meta charset="UTF-8">
				<title>Search for Item</title>
			</head>
			<body>
				<h1>Search for Item</h1>
				<form method="POST">
					<label>Item Name:</label>
					<input type="text" name="name" required>
					<button type="submit">Search</button>
				</form>
			</body>
			</html>
		`)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	itemName := r.FormValue("name")

	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM store WHERE name LIKE $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query("%" + itemName + "%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Raiting, &item.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, items)
	CheckError(err)
}

// func giveRating(w http.ResponseWriter, r *http.Request) {

// 	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, ps, dbname)
// 	db, err = sql.Open("postgres", db_connection)
// 	CheckError(err)
// 	defer db.Close()

// 	id := r.FormValue("id")
// 	newValue := r.FormValue("new_value")

// 	if (id == "" || newValue == "" ){
// 		fmt.Fprintf(w, "ID or value can not be empty")
// 	} else  {

// 		stmt, err := db.Prepare("UPDATE store SET rating = (rating + $1)/2 WHERE id = $2")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		defer stmt.Close()

// 		_, err = stmt.Exec(newValue, id)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// 	http.Redirect(w, r, "templates/showall.html", http.StatusSeeOther)
// }
