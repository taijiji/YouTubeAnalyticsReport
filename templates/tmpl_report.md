# show int レポート

## 動画名

{{ . }}
{{ range .Video_list }}
1. [{{ .Video_title }}](https://www.youtube.com/watch?v={{ .Video_id }}) ( {{ .Updated_date }} 公開)
{{ end }}

{{ range .param }}
|||
|---|---|
|動画名|{{ .param.Video_title }}|
|動画URL|https://www.youtube.com/watch?v={{ .param.Video_id }}|
|動画公開日|{{ .param.Updated_date }}|
|集計期間||
|サムネイル|<img src="images/thumbnail_{{ .param.Video_id }}.jpg">|
|再生回数|{{ .param.View_counts }}|
|グッド回数|{{ .param.Like_counts }}|
|バッド回数|{{ .param.Dislike_counts }}|
|||
<div style="page-break-before:always"></div>
{{ end }}
