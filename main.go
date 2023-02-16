package main

import (
	// sr "MyStore/structs.go"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"time"
	"unicode"
	// _ "github.com/lib/pq"
)

type User struct {
	Username string
	Mail     string
	Password string
}

type Item struct {
	ID    int64
	Name  string
	Price float32
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "db_go"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	for true {
		clearScreen()
		var choice string
		Hello()
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			Register()
		case "2":
			Login()
		case "3":
			bye()
			os.Exit(1)
		default:
			fmt.Printf("command not defined, please choose again")
		}
		time.Sleep(time.Second)
	}
}

func UserInit(nick, pswrd string) {
	
}

func Login() {
	var nick, password string
	fmt.Println("Enter your nickname:  ")
	fmt.Scanln(&nick)
	fmt.Println("Enter your password:  ")
	fmt.Scanln(&password)
	if searchUser(nick, password) {
		clearScreen()
		fmt.Println("Logged in successfuly")
		Shop()
	} else {
		clearScreen()
		fmt.Println("Invalid nickname or password, try again or create new account")
	}
}
func Shop() {
	clearScreen()
	fmt.Println("********************************************************************************************************")
	fmt.Println("***********************************      Welcome to out Store      *************************************")
	fmt.Println("	Enter 1 to show all items")
	fmt.Println("	Enter 2 to filter items by price")
	fmt.Println("	Enter 3 to filter items by rating")
	fmt.Println("	Enter 4 to search item by name")
	fmt.Println("	Enter 5 to go back")

	var cmnd string
	fmt.Scanln(&cmnd)
	switch cmnd {
	case "1":
		showAllItems()
	case "2":
		filterItemsByPrice()
	case "3":
		filterItemsByRating()
	case "4":
		searchItemByName()
	case "5":
		return
	default:
		fmt.Println("command not defined, please choose again")
		time.Sleep(2 * time.Second)
		Shop()
	}
}
func searchItemByName() {
	fmt.Println("Enter item name: ")
	var find string
	fmt.Scanln(&find)
	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where quantity > 0 and name = $1`, find)
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
		bye()
		os.Exit(1)
	}
}
func GiveRating(name string) {
	fmt.Println("Rate an item from 0 to 10")
	var nw float32
	fmt.Scanln(&nw)
	rows, err := db.Query(`SELECT rating FROM lineitem where name = $1`, name)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var raiting float32
		err = rows.Scan(&raiting)
		CheckError(err)
		if raiting == 0 {
			raiting = nw
		} else {
			raiting = (raiting + nw) / 2
		}

		_, err = db.Exec(`update lineitem set rating = $1 where name = $2`, raiting, name)
		CheckError(err)
		// clearScreen()
		fmt.Println("Rating added....")
		time.Sleep(2 * time.Second)
		Shop()
	}
}
func filterItemsByRating() {
	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where quantity > 0 order by rating`)
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
		// clearScreen()
		Shop()
	default:
		bye()
		os.Exit(1)
	}
}
func filterItemsByPrice() {
	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where quantity > 0 order by price`)
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
		// clearScreen()
		Shop()
	default:
		bye()
		os.Exit(1)
	}
}

func showAllItems() {

	rows, err := db.Query(`SELECT name, price, rating FROM lineitem where quantity > 0`)
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
		// clearScreen()
		Shop()
	default:
		bye()
		os.Exit(1)
	}
}
func Register() {
	var nick string
	var pswrd string
	createNick(&nick)
	createPassword(&pswrd)

	newUser := new(User)
	// newUser.UserInit(nick, pswrd)

	_, err := db.Exec(insertUserStat(), newUser.Mail, newUser.Password)
	CheckError(err)
	// clearScreen()
	fmt.Println("Registration was successful, you will go back to login")
}
func searchUser(nickname string, pswrd string) bool {
	rows, err := db.Query(`SELECT nickname, pswrd FROM users where nickname = $1 and pswrd = $2`, nickname, pswrd)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var nick string
		var psw string

		err = rows.Scan(&nick, &psw)
		CheckError(err)

		if nick != "" {
			time.Sleep(2 * time.Second)
			return true
		}
	}
	CheckError(err)
	return false
}
func insertUserStat() string {
	return `
INSERT INTO users (nickname, pswrd)
VALUES ($1, $2)`
}
func selectNick() string {
	return `
	select nick from users
	where nick = $1`
}
func createNick(n *string) {
	fmt.Print("Create nickname:  ")
	fmt.Scanln(n)
	rows, err := db.Query(`SELECT nickname FROM users where nickname = $1`, n)
	CheckError(err)
	for rows.Next() {
		var nick string
		err = rows.Scan(&nick)
		CheckError(err)

		if nick == *n {
			fmt.Println("naickname already exist, please create another nickname")
			time.Sleep(time.Second * 2)
			createNick(n)
		}
	}

}
func createPassword(pswrd *string) {
	fmt.Print("Create password:\n\n")
	fmt.Println("Password should have:")
	fmt.Println("at least lenth 5")
	fmt.Println("at least one uppercase letter")
	fmt.Println("at least one lowercase letter")
	fmt.Println("at least one digit")

	fmt.Scanln(pswrd)
	if validPassword(*pswrd) {
		fmt.Println("Password created")
		time.Sleep(time.Second * 2)
	} else {
		fmt.Println("Invalid password, please enter new one")
		time.Sleep(time.Second * 2)
		createPassword(pswrd)
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

func Hello() {
	clearScreen()
	fmt.Println("********************************************************************************************************")
	fmt.Println("*******************************************      MENU      *********************************************")
	fmt.Println("	Enter 1 to Register")
	fmt.Println("	Enter 2 to Login")
	fmt.Println("	Enter 3 to Exit")
}
func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func bye() {
	clearScreen()
	fmt.Print("********************************************************************************************************\n\n")
	fmt.Println("					         GOODBYE")
	fmt.Println("\n********************************************************************************************************")
}
