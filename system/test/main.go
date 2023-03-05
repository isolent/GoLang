package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// User struct to store user information
type User struct {
	Username string `json:"username"`
	Mail     string `json:"email"`
	Password string `json:"password"`
}

// Item struct to store item information
type Item struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
}

// Store struct to store all data
type Store struct {
	Users []User `json:"users"`
	Items []Item `json:"items"`
}

func (s *Store) Register(username, email, password string) error {
	// read the existing JSON data from the file
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			// create an empty slice of User structs if the file does not exist
			s.Users = []User{}
		} else {
			return err
		}
	} else {
		// deserialize the existing JSON data into the Users slice
		err = json.Unmarshal(jsonData, &s.Users)
		if err != nil {
			return err
		}
	}

	// add the new User to the Users slice
	user := User{username, email, password}
	s.Users = append(s.Users, user)

	// serialize the Users slice back into JSON data
	jsonData, err = json.Marshal(s.Users)
	if err != nil {
		return err
	}

	// write the JSON data back to the file
	err = ioutil.WriteFile("users.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Authorise(mail, password string) bool {
	// read the JSON data from the file
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		return false
	}

	// deserialize the JSON data into a slice of User structs
	err = json.Unmarshal(jsonData, &s.Users)
	if err != nil {
		return false
	}

	// check if the provided username and password match with any of the Users
	for _, user := range s.Users {
		if user.Mail == mail && user.Password == password {
			return true
		}
	}

	return false
}

func (s *Store) Search(name string) []Item {
	// read the JSON data from the file
	jsonData, err := ioutil.ReadFile("storeitems.json")
	if err != nil {
		return nil
	}

	// deserialize the JSON data into a slice of Item structs
	err = json.Unmarshal(jsonData, &s.Items)
	if err != nil {
		return nil
	}

	var results []Item
	// search for Items with the provided name
	for _, item := range s.Items {
		if item.Name == name {
			results = append(results, item)
		}
	}

	return results
}

func (s *Store) Filter(price float64, rating float64) []Item {
	// read the JSON data from the fileSS
	jsonData, err := ioutil.ReadFile("storeitems.json")
	if err != nil {
		return nil
	}

	// deserialize the JSON data into a slice of Item structs
	err = json.Unmarshal(jsonData, &s.Items)
	if err != nil {
		return nil
	}

	var results []Item
	// filter the Items based on the provided price and rating parameters
	for _, item := range s.Items {
		if item.Price <= price && item.Rating >= rating {
			results = append(results, item)
		}
	}

	return results
}

func (s *Store) GiveRating(name string, rating float64) error {
	// read the JSON data from the file
	jsonData, err := ioutil.ReadFile("storeitems.json")
	if err != nil {
		return err
	}

	// deserialize the JSON data into a slice of Item structs
	err = json.Unmarshal(jsonData, &s.Items)
	if err != nil {
		return err
	}

	// find the Item with the provided name and give it the provided rating
	found := false
	for i, item := range s.Items {
		if item.Name == name {
			s.Items[i].Rating = rating
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Item not found")
	}

	// serialize the modified slice of Item structs into JSON data
	jsonData, err = json.MarshalIndent(s.Items, "", "  ")
	if err != nil {
		return err
	}

	// write the JSON data back to the file
	err = ioutil.WriteFile("storeitems.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var store Store

	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Search items")
	fmt.Println("4. Filter items")
	fmt.Println("5. Give item rating")
	fmt.Print("Choose an option: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter name: ")
		var name string
		fmt.Scan(&name)

		fmt.Print("Enter email: ")
		var email string
		fmt.Scan(&email)

		fmt.Print("Enter password: ")
		var password string
		fmt.Scan(&password)

		store.Register(name, email, password)
	case 2:
		fmt.Print("Enter email: ")
		var email string
		fmt.Scan(&email)

		fmt.Print("Enter password: ")
		var password string
		fmt.Scan(&password)

		if store.Authorise(email, password) {
			fmt.Println("Authorised.")
		} else {
			fmt.Println("Unauthorised.")
		}
	case 3:
		fmt.Print("Enter item name: ")
		var name string
		fmt.Scan(&name)

		items := store.Search(name)
		for _, item := range items {
			fmt.Println("Name:", item.Name)
			fmt.Println("Price:", item.Price)
			fmt.Println("Rating:", item.Rating)
		}
	case 4:
		fmt.Print("Enter maximum price: ")
		var price float64
		fmt.Scan(&price)

		fmt.Print("Enter minimum rating: ")
		var rating float64
		fmt.Scan(&rating)

		items := store.Filter(price, rating)
		for _, item := range items {
			fmt.Println("Name:", item.Name)
			fmt.Println("Price:", item.Price)
			fmt.Println("Rating:", item.Rating)
		}
	case 5:
		fmt.Print("Enter item name: ")
		var name string
		fmt.Scan(&name)

		fmt.Print("Enter rating: ")
		var rating float64
		fmt.Scan(&rating)

		store.GiveRating(name, rating)
	}
}
