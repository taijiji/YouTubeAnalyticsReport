# show int 年間レポート

## 動画名
|サムネイル|動画名|公開日|再生回数|
|---|---|---|---|
{{ range . -}}
|<img src="images/thumbnail_{{ .Video_id }}_trim.jpg">|[{{ .Video_title }}](https://www.youtube.com/watch?v={{ .Video_id }})|{{ .Updated_date }}|{{ .View_counts }}|
{{ end -}}