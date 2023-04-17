

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	unsplashAPIBaseURL = "https://api.unsplash.com"
	unsplashRandomPath = "/photos/random"
)

type unsplashResponse struct {
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg := &sync.WaitGroup{}
	
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("6275836826:AAFkpwEdEnZa6YDcGR0TxFUtCjn7ySO4K4k", opts...)
	if err != nil {
		panic(err)
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/image", bot.MatchTypeExact, handler)

	b.Start(ctx)

	


	wg.Wait()
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	imageURL, err := getRandomUnsplashImageURL()
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "DIDN'T GET IT",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   imageURL,
	})
}

func getRandomUnsplashImageURL() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, unsplashAPIBaseURL+unsplashRandomPath, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Client-ID iNm_s6QlxiEZdM4h4BK2GRBHbMCpGHpQmANRePORg3g")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var unsplashResp unsplashResponse
	err = json.NewDecoder(resp.Body).Decode(&unsplashResp)
	if err != nil {
		return "", err
	}
	return unsplashResp.URLs.Regular, nil
}

