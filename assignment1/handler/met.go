package handler

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"time"
	"unicode"

	_ "github.com/lib/pq"
)

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

func Welcome() {
	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", db_connection)
	CheckError(err)
	defer db.Close()

	fmt.Println("Register or Login:")
	fmt.Println("Choose 1: Register")
	fmt.Println("Choose 2: Login")
	fmt.Println("Choose 3: Exit")

	{
		var choice string
		fmt.Scanln(&choice)
		if choice == "1" {
			Register()
		} else if choice == "2" {
			Login()
		} else if choice == "3" {
			theEnd()
			os.Exit(1)
		} else {
			fmt.Printf("Command is not found, try again")
		}
		time.Sleep(time.Second)
	}
}

func Register() {
	var nick string
	var password string
	createNickname(&nick)
	createPassword(&password)

	newUser := new(User)
	newUser.UserInit(nick, password)

	_, err := db.Exec(insertUser(), newUser.name, newUser.Password)
	CheckError(err)
	fmt.Println("Successfully registered")
}

func Login() {
	var nick, password string
	fmt.Println("Enter your name:  ")
	fmt.Scanln(&nick)
	fmt.Println("Enter your password:  ")
	fmt.Scanln(&password)
	if searchUser(nick, password) {
		fmt.Println("\nLogged in successfully")
		Shop()
	} else {
		fmt.Println("Invalid name or password, try again or create new account")
		Welcome()
	}
}
func Shop() {

	fmt.Println("What you want to do:")
	fmt.Println("Choose 1: Show all items")
	fmt.Println("Choose 2: Filter items by price")
	fmt.Println("Choose 3: Filter items by rating")
	fmt.Println("Choose 4: Search item by name")
	fmt.Println("Choose 5: Exit")

	var cmnd string
	fmt.Scanln(&cmnd)
	if cmnd == "1" {
		showAllItems()
	} else if cmnd == "2" {
		filterItemsByPrice()
	} else if cmnd == "3" {
		filterItemsByRating()
	} else if cmnd == "4" {
		searchItemByName()
	} else if cmnd == "5" {
		Welcome()
	} else {
		fmt.Println("Command is not found, try again")
		Shop()
	}
}

func showAllItems() {

	rows, err := db.Query(`SELECT name, price, rating FROM store where quantity > 0`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var price float32
		var raiting float32

		err = rows.Scan(&name, &price, &raiting)
		CheckError(err)
		fmt.Printf(" %-15s  %-15f %-15f \n", name, price, raiting)
	}
	fmt.Println("\nEnter 1 to go back")
	fmt.Println("\nEnter anything else to exit")
	var cmnd string
	fmt.Scanln(&cmnd)
	switch cmnd {
	case "1":
		Shop()
	default:
		theEnd()
		os.Exit(1)
	}
}

func filterItemsByPrice() {
	rows, err := db.Query(`SELECT name, price, rating FROM store where quantity > 0 order by price`)
	CheckError(err)

	defer rows.Close()
	fmt.Printf(" %-15s  %-15s %-15s \n", "name", "price", "raiting")
	for rows.Next() {
		var name string
		var price float32
		var raiting float32

		err = rows.Scan(&name, &price, &raiting)
		CheckError(err)

		fmt.Printf(" %-15s  %-15f %-15f \n", name, price, raiting)
	}
	fmt.Println("\nEnter 1 to go back")
	fmt.Println("\nEnter anything else to exit")
	var cmnd string
	fmt.Scanln(&cmnd)
	if cmnd == "1" {
		Shop()
	} else {
		theEnd()
		os.Exit(1)
	}
}

func filterItemsByRating() {
	rows, err := db.Query(`SELECT name, price, rating FROM store where quantity > 0 order by rating`)
	CheckError(err)

	defer rows.Close()
	fmt.Printf(" %-15s  %-15s %-15s \n", "name", "price", "raiting")
	for rows.Next() {
		var name string
		var price float32
		var raiting float32

		err = rows.Scan(&name, &price, &raiting)
		CheckError(err)

		fmt.Printf(" %-15s  %-15f %-15f \n", name, price, raiting)
	}
	fmt.Println("\nEnter 1 to go back")
	fmt.Println("\nEnter anything else to exit")
	var cmnd string
	fmt.Scanln(&cmnd)
	switch cmnd {
	case "1":
		Shop()
	default:
		theEnd()
		os.Exit(1)
	}
}

func searchItemByName() {
	fmt.Println("Enter item name: ")
	var find string
	fmt.Scanln(&find)
	rows, err := db.Query(`SELECT name, price, rating FROM store where quantity > 0 and name = $1`, find)
	CheckError(err)

	defer rows.Close()

	noRow := false
	for rows.Next() {
		var name string
		var price float32
		var raiting float32

		err = rows.Scan(&name, &price, &raiting)
		CheckError(err)
		fmt.Printf(" %-15s  %-15s %-15s \n", "name", "price", "raiting")
		fmt.Printf(" %-15s  %-15f %-15f \n", name, price, raiting)
		noRow = true
	}
	if !noRow {
		fmt.Println("\nNothing found, you will go back")
		time.Sleep(2 * time.Second)
		Shop()
	}
	fmt.Println("\nEnter 1 to go back")
	fmt.Println("Enter 2 to give rating")
	fmt.Println("Enter anything else to exit")
	var cmnd string
	fmt.Scanln(&cmnd)
	switch cmnd {
	case "1":
		Shop()
	case "2":
		GiveRating(find)
	default:
		theEnd()
		os.Exit(1)
	}
}
func GiveRating(name string) {
	fmt.Println("Rate an item from 0 to 10")
	var neww float32
	fmt.Scanln(&neww)
	rows, err := db.Query(`SELECT rating FROM store where name = $1`, name)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var raiting float32
		err = rows.Scan(&raiting)
		CheckError(err)
		if raiting == 0 {
			raiting = neww
		} else {
			raiting = (raiting + neww) / 2
		}

		_, err = db.Exec(`update store set rating = $1 where name = $2`, raiting, name)
		CheckError(err)
		fmt.Println("Rating added")
		time.Sleep(2 * time.Second)
		Shop()
	}
}

func searchUser(name string, password string) bool {
	rows, err := db.Query(`SELECT name, password FROM users where name = $1 and password = $2`, name, password)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var nick string
		var psw string

		err = rows.Scan(&nick, &psw)
		CheckError(err)

		if nick != "" {
			return true
		}
	}
	CheckError(err)
	return false
}

func insertUser() string {
	return `INSERT INTO users (name, password) VALUES ($1, $2)`
}

func createNickname(n *string) {
	fmt.Print("Create name:  ")
	fmt.Scanln(n)
	rows, err := db.Query(`SELECT name FROM users where name = $1`, n)
	CheckError(err)
	for rows.Next() {
		var nick string
		err = rows.Scan(&nick)
		CheckError(err)

		if nick == *n {
			fmt.Println("name already exist, please create another name")
			time.Sleep(time.Second * 2)
			createNickname(n)
		}
	}

}
func createPassword(password *string) {
	fmt.Print("Create password:\n\n")
	fmt.Println("Password should have:")
	fmt.Println("at least length 5")
	fmt.Println("at least one uppercase letter")
	fmt.Println("at least one lowercase letter")
	fmt.Println("at least one digit")

	fmt.Scanln(password)
	if validPassword(*password) {
		fmt.Println("Password created")
		time.Sleep(time.Second * 2)
	} else {
		fmt.Println("Invalid password, please enter new one")
		time.Sleep(time.Second * 2)
		createPassword(password)
	}
}

func validPassword(p string) bool {
	var (
		hasLen   = false
		hasUpper = false
		hasLower = false
		hasDigit = false
	)
	if len(p) >= 5 {
		hasLen = true
	}
	for _, v := range p {
		switch {
		case unicode.IsUpper(v):
			hasUpper = true
		case unicode.IsLower(v):
			hasLower = true
		case unicode.IsDigit(v):
			hasDigit = true
		}
	}
	return hasLen && hasDigit && hasLower && hasUpper
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func theEnd() {
	clearScreen()
	fmt.Println("Session is ended!")
}

type User struct {
	name     string
	Password string
}

func (u *User) UserInit(L string, P string) {
	u.name = L
	u.Password = P
}
