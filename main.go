package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	channel_stats := getChannelStats()
	startdate := "2023-03-06T00:00:00Z"
	enddate := "2023-05-08T23:00:00Z"

	video_stats := getVideoStats(startdate, enddate)

	fmt.Println(channel_stats)
	fmt.Println(video_stats)
}
