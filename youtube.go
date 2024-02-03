package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
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
	thumbnail_url  string
}

func reverseVideoList(v []Video) []Video {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
	return v
}

func downloadImage(url string, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("Image downloaded successfully : " + filename)
	return nil

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
		//channel_stats.Channel_id = ch.Id
		//fmt.Println(strings.Repeat("=", 50))
		//fmt.Println("Channel Title: ", channel.Snippet.Title)
		//fmt.Println("Channel ID: ", channel.Id)
		//fmt.Printf("Suscribers: %v\n", channel.Statistics.SubscriberCount)
		//fmt.Println(strings.Repeat("=", 50))
	}
	return channel_stats
}

// Get Statics Data from YouTube Analytics API
func getVideoAnalytics(videoId string) {

	ctx := context.Background()
	//channelID := os.Getenv("CHANNEL_ID")

	// Google Cloud Platform の OAuth 2.0 クライアント ID と秘密鍵を取得する
	//clientID := os.Getenv("GOOGLE_CLIENT_ID")
	//clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	//clientID := "1012177613960-l6ududdcael66n0ph7squ836tpd542oc.apps.googleusercontent.com"
	//clientSecret := "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"
	//
	//// OAuth 2.0 トークンを取得する
	//config := oauth2.Config{
	//	ClientID:     clientID,
	//	ClientSecret: clientSecret,
	//	Endpoint: oauth2.Endpoint{
	//		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	//		TokenURL: "https://oauth2.googleapis.com/token",
	//	},
	//}

	b, err := os.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	scope := youtube.YoutubeReadonlyScope
	config, err := google.ConfigFromJSON(b, scope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	fmt.Println("test00")
	fmt.Println(config)

	config.RedirectURL = "http://localhost:8090"

	// 以下 oauth2.goからのコピー。
	// cacheFile, err := tokenCacheFile()
	// if err != nil {
	// 	log.Fatalf("Unable to get path to cached credential file. %v", err)
	// }
	// token, err := tokenFromFile(cacheFile)
	// if err != nil {
	// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	// 	if launchWebServer {
	// 		fmt.Println("Trying to get token from web")
	// 		token, err = getTokenFromWeb(config, authURL)
	// 	} else {
	// 		fmt.Println("Trying to get token from prompt")
	// 		token, err = getTokenFromPrompt(config, authURL)
	// 	}
	// 	if err == nil {
	// 		saveToken(cacheFile, tok)
	// 	}
	// }

	//token, err := config.Exchange(ctx, os.Getenv("API_KEY"))

	////////////////////  最新 update 2023 /12/18 ///////////////////////////////
	// ここでエラーでてる。
	// oauth2: cannot fetch token: 400 Bad Request
	// Response: {
	//   "error": "invalid_grant",
	//   "error_description": "Malformed auth code."
	// }
	token, err := config.Exchange(ctx, "1012177613960-l6ududdcael66n0ph7squ836tpd542oc.apps.googleusercontent.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("test01")

	//client := getClient(youtubeanalytics.YtAnalyticsReadonlyScope)
	//scope := "https://www.googleapis.com/auth/yt-analytics.readonly"
	//client := getClient(scope)
	//service, err := youtube.New(client)
	service, err := youtube.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))

	fmt.Println("test1")

	//// チャンネルの統計情報を取得します。
	//request := service.Channels.List([]string{"snippet", "contentDetails", "statistics"}).Id(channelID)
	//response, err := request.Do()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 視聴者の年齢層を取得します。
	//for _, item := range response.Items {
	//	audience := item.Statistics.Audience
	//	ageRanges := audience.AgeRanges
	//	for _, ageRange := range ageRanges {
	//		fmt.Println("年齢層:", ageRange.Name, ", 割合:", ageRange.Percentage)
	//	}
	//}

	// 視聴者の年齢層のデータを取得します。
	response, err := service.Videos.List([]string{"snippet", "contentDetails", "statistics", "topicDetails", "viewerStats"}).Id(videoId).Do()
	if err != nil {
		panic(err)
	}
	fmt.Print(response)

	fmt.Println("test2")

	//for _, item := range response.Items {
	//	fmt.Println("視聴者の年齢層:")
	//	fmt.Println("  - 18-24歳:", item.)
	//	fmt.Println("  - 25-34歳:", item.ViewerStats.AgeGroup25To34)
	//	fmt.Println("  - 35-44歳:", item.ViewerStats.AgeGroup35To44)
	//	fmt.Println("  - 45-54歳:", item.ViewerStats.AgeGroup45To54)
	//	fmt.Println("  - 55-64歳:", item.ViewerStats.AgeGroup55To64)
	//	fmt.Println("  - 65歳以上:", item.ViewerStats.AgeGroup65Plus)
	//
	//}

}

// Get Statics Data from YouTube Data API
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
		video.thumbnail_url = search.Snippet.Thumbnails.High.Url

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

		image_name := "reports/images/thumbnail_" + video.Video_id + ".jpg"
		downloadImage(video.thumbnail_url, image_name)
	}
	video_list = reverseVideoList(video_list)
	return video_list
}
