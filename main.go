package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	call_channel := service.Channels.List([]string{"snippet", "statistics"}).Id(CHANNEL_ID)
	response, err := call_channel.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	for _, channel := range response.Items {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Channel Title: ", channel.Snippet.Title)
		fmt.Println("Channel ID: ", channel.Id)
		fmt.Printf("Suscribers: %v\n", channel.Statistics.SubscriberCount)
		fmt.Println(strings.Repeat("=", 50))
	}

	call_search := service.Search.List([]string{"id", "snippet"}).ChannelId(CHANNEL_ID).Order("date").MaxResults(15)
	response_search, err_search := call_search.Do()
	if err_search != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	for _, search := range response_search.Items {
		fmt.Println(strings.Repeat("-", 100))
		fmt.Println("Video Title: ", search.Snippet.Title)
		fmt.Println("Video ID: ", search.Id)
		fmt.Println("Uploaded Date: ", search.Snippet.PublishedAt)
	}
}
