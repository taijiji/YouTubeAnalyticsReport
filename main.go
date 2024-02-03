package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//channel_stats := getChannelStats()
	//startdate := "2023-07-30T00:00:00Z"
	//enddate := "2023-10-03T23:00:00Z"

	//today := time.Now().Format("20060102")
	//video_list := getVideoStats(startdate, enddate)
	video_id := "Kv_EaxHzprk"
	getVideoAnalytics(video_id)

	//render_report(video_list, today)
}
