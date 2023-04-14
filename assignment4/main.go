package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Set random seed for image selection
	rand.Seed(time.Now().UnixNano())

	// Listen for updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// Check if the update contains a message
		if update.Message == nil {
			continue
		}

		// Check if the message is "image" or "/image"
		if update.Message.Text == "image" || update.Message.Command() == "image" {
			// Get a random image from Unsplash API
			imageURL := getRandomImage()

			// Send the image as a photo
			photo := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, imageURL)
			bot.Send(photo)
		}
	}
}

func getRandomImage() string {
	// Make a request to Unsplash API to get a random photo
	resp, err := http.Get("https://api.unsplash.com/photos/random?client_id=" + os.Getenv("UNSPLASH_ACCESS_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the response to get the image URL
	var image map[string]interface{}
	err = json.Unmarshal(body, &image)
	if err != nil {
		log.Fatal(err)
	}

	return image["urls"].(map[string]interface{})["regular"].(string)
}