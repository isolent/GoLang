package main

import (
	"encoding/json"
	"fmt"
	_ "math/rand"
	"net/http"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	unsplashAPIURL = "https://api.unsplash.com/photos/random"
	unsplashAPIKey = "iNm_s6QlxiEZdM4h4BK2GRBHbMCpGHpQmANRePORg3g"
)

var (
	counter      int
	counterMutex sync.Mutex //безопасный доступ к переменной каунтер
)

func main() {
	// инициализация телеграм бота с помощью токена
	bot, err := tgbotapi.NewBotAPI("6275836826:AAFkpwEdEnZa6YDcGR0TxFUtCjn7ySO4K4k")
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "image" || update.Message.Command() == "image" {
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				incrementCounter()
				wg.Done()
			}()
			wg.Wait()

			photo, err := getRandomPhoto()
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, I couldn't find a random photo."))
				continue
			}

			fmt.Println(photo.URLs.Regular)
			msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, photo.URLs.Regular)
			msg.Caption = fmt.Sprintf("Image with id: %d", counter)
			bot.Send(msg)
		}
	}
}

func incrementCounter() {
	counterMutex.Lock()
	counter++
	counterMutex.Unlock()
}

func getRandomPhoto() (*UnsplashPhoto, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, unsplashAPIURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Version", "v1")
	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", unsplashAPIKey))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var photo UnsplashPhoto
	err = json.NewDecoder(resp.Body).Decode(&photo)
	if err != nil {
		return nil, err
	}

	return &photo, nil
}

type UnsplashPhoto struct {
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}
