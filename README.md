# YouTubeAnalyticsReport
The script will get videos for an arbitrarily designated time period　on your YouTube Channel.

# Generated report sample

See [ Report Sample](./reports/showint_report_20231021.md)

# How to run this script

Make `.env` file with your YouTube API KEY(YouTube Data API v3) & Channel ID.
```
vi .env

API_KEY="XXXXXXXX"
CHANNEL_ID="XXXXXXXX"
```

Update `startdate` & `enddate` on [main.go](main.go) like this

```
	startdate := "2023-05-15T00:00:00Z"
	enddate := "2023-07-24T23:00:00Z"
```

Run the scprits like this. 

```
$ go run main.go oauth2.go report.go youtube.go image.go
```