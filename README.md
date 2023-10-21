# YouTubeAnalyticsReport
Add your [YouTube API key](https://developers.google.com/youtube/registering_an_application) and your YouTube Channel ID into `.env` file.

# Generated report sample

See [ Report Sample](./reports/showint_report_20230902.md)

# How to run this script

Make `.env` file with your YouTube API KEY(YouTube Data API v3) & Channel ID.
```
vi .env

API_KEY="XXXXXXXX"
CHANNEL_ID="XXXXXXXX"
```

Update start date & end date on [main.go](main.go) like this

```
	startdate := "2023-05-15T00:00:00Z"
	enddate := "2023-07-24T23:00:00Z"
```

Run the scprits like this. The script will get videos for the period on your YouTube Channel.

```
$ go run main.go oauth2.go report.go youtube.go
```