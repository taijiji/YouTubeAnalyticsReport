# show int レポート

## 動画名

{{ range . }}
1. [{{ .Video_title }}](https://www.youtube.com/watch?v={{ .Video_id }}) ( {{ .Updated_date }} 公開)
{{ end }}

{{ range . }}
|||
|---|---|
|動画名|{{ .Video_title }}|
|動画公開日|{{ .Updated_date }}|
|サムネイル|<img src="images/thumbnail_{{ .Video_id }}.jpg">|
|再生回数|{{ .View_counts }}|
|グッド回数|{{ .Like_counts }}|
|バッド回数|{{ .Dislike_counts }}|
|||
<div style="page-break-before:always"></div>
{{ end }}
