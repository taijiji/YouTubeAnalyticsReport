package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//channel_stats := getChannelStats()
	startdate := "2024-01-01T00:00:00Z"
	enddate := "2024-02-04T23:00:00Z"

	today := time.Now().Format("20060102")
	video_list := getVideoStats(startdate, enddate)

	render_report(video_list, today)
}
