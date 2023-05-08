package main

import (
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type ChannelStats struct {
	Channel_title string
	Subscribers   string
	Channel_id    string
}

type Video struct {
	Video_title    string
	Video_id       string
	Updated_date   string
	View_counts    uint64
	Like_counts    uint64
	Dislike_counts uint64
}

func reverseVideoList(v []Video) []Video {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
	return v
}

func getChannelStats() ChannelStats {
	API_KEY := os.Getenv("API_KEY")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")

	var channel_stats ChannelStats

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(API_KEY))

	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	call_channel := service.Channels.List([]string{"snippet", "statistics"}).Id(CHANNEL_ID)
	response_channel, err := call_channel.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	for _, ch := range response_channel.Items {
		channel_stats.Channel_title = ch.Snippet.Title
		channel_stats.Subscribers = string(ch.Statistics.SubscriberCount)
		channel_stats.Channel_id = ch.Id
		//fmt.Println(strings.Repeat("=", 50))
		//fmt.Println("Channel Title: ", channel.Snippet.Title)
		//fmt.Println("Channel ID: ", channel.Id)
		//fmt.Printf("Suscribers: %v\n", channel.Statistics.SubscriberCount)
		//fmt.Println(strings.Repeat("=", 50))
	}
	return channel_stats
}

func getVideoStats(startdate string, enddate string) []Video {

	API_KEY := os.Getenv("API_KEY")
	CHANNEL_ID := os.Getenv("CHANNEL_ID")

	var video Video
	var video_list []Video

	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	//This is draft code for YouTubeAnalitycs API with Oauth2.0
	//client := getClient(youtube.YoutubeReadonlyScope)
	//service, err := youtube.New(client)

	call_search := service.Search.List([]string{"id", "snippet"}).ChannelId(CHANNEL_ID).Order("date").MaxResults(15).PublishedAfter(startdate).PublishedBefore(enddate)
	response_search, err_search := call_search.Do()
	if err_search != nil {
		log.Fatalf("Error making search API call: %v", err_search)
	}

	for _, search := range response_search.Items {
		video.Video_title = search.Snippet.Title
		video.Video_id = search.Id.VideoId
		video.Updated_date = strings.Split(search.Snippet.PublishedAt, "T")[0]

		call_video := service.Videos.List([]string{"snippet", "contentDetails", "statistics"}).Id(search.Id.VideoId)
		response_video, err := call_video.Do()
		if err != nil {
			log.Fatalf("Error making Video API call: %v", err)
		}
		for _, video_stats := range response_video.Items {
			video.View_counts = video_stats.Statistics.ViewCount
			video.Like_counts = video_stats.Statistics.LikeCount
			video.Dislike_counts = video_stats.Statistics.DislikeCount
		}

		// 公式ページには、Go言語でYouTube Analytics APIを利用できるサンプルが無かった。
		// インプレッション数や視聴者の性別、トラフィック流入元などのデータを取るには Analytics APIへのアクセスが必須。
		// ふと見つけた google.golang.org/api/youtubeanalytics/v2 が利用可能かもしれない。
		// https://pkg.go.dev/google.golang.org/api/youtubeanalytics/v2#pkg-functions
		// ただし公開日が2023/1/23。まだReadyじゃない？いちおうGoogleの正式なプロダクトに見える。
		// まだ正しい使い方がわからないので、要検討。だめなら、Pythonで書き直す必要があるかも。
		// https://developers.google.com/youtube/reporting/v1/code_samples/python
		// 実現できる手段が無いわけではないが、Go言語で実現できるかは微妙。要検証。
		// この資料、読んで見る価値あるかも
		// https://blog.codecamp.jp/youtube-analytics-python

		//TODO: This is draft code for YouTubeAnalitycs API with Oauth2.0
		//client_analytics := getClient(youtubeanalytics.YtAnalyticsReadonlyScope)
		//service_analytics, err := youtubeanalytics.New(client_analytics)
		//if err != nil {
		//	log.Fatalf("Error creating new YouTube service: %v", err)
		//}
		// 動作OK
		//call_analytics := service_analytics.Reports.Query().Ids("channel==MINE").Dimensions("ageGroup").Metrics("viewerPercentage").StartDate("2023-01-01").EndDate("2023-02-11")

		// 動作OK
		//call_analytics := service_analytics.Reports.Query().Ids("channel==MINE").Dimensions("channel").Metrics("views,likes").StartDate("2023-01-01").EndDate("2023-02-11")

		// 動作OK
		//call_analytics := service_analytics.Reports.Query().Ids("channel==MINE").Dimensions("day").Metrics("views,likes").StartDate("2023-01-01").EndDate("2023-02-11")

		// Ids("channel==MINE") は必須。
		// Dimensions("video") がなぜが使えない。
		// Dimensions("day")やDimensions("ageGroup")、Dimensions("channel")は使えてる。
		// 2023/03/03 18:41:27 Error making Analytics API call: googleapi: Error 400: The query is not supported. Check the documentation at https://developers.google.com/youtube/analytics/v2/available_reports for a list of supported queries., badRequest

		// 動いた！ .Sort("-views")が必須ぽい。
		// annotationImpressions は結果が0になる。これはAPI server側ががまだ対応していないのかもしれない。なんかいろいろ制約があるようにみえるな。
		// Note: YouTube Analytics API reports only return data for the annotationClickThroughRate and annotationCloseRate metrics as of June 10, 2012. In addition, YouTube Analytics API reports only return data for the remaining annotation metrics as of July 16, 2013.
		// 動作OK

		//TODO: This is draft code for YouTubeAnalitycs API with Oauth2.0
		//call_analytics := service_analytics.Reports.Query().Ids("channel==MINE").Dimensions("video").Metrics("views,likes,//////annotationImpressions").StartDate("2022-01-01").EndDate("2023-02-11").MaxResults(10).Sort("-views")

		//response_anlytics, err := call_analytics.Do()
		//if err != nil {
		//	log.Fatalf("Error making Analytics API call: %v", err)
		//}

		video_list = append(video_list, video)
	}
	video_list = reverseVideoList(video_list)
	return video_list
}
