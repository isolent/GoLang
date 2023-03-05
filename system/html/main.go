package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "database/sql"
	// _ "github.com/lib/pq"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"bob", "asd", "1234"}
	// fmt.Fprintf(w, `<b>Main page </b>`)
	temp, _ := template.ParseFiles("templates/home_page.html")
	temp.Execute(w, bob)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("UserName is: %s. His Email: %s. Password: %s", u.Name, u.Email, u.Password)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")

}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":4334", nil)
}
func main() {
	handleRequest()
}
