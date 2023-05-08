# show int レポート

## 動画名

{{ range . }}
1. [{{ .Video_title }}](https://www.youtube.com/watch?v={{ .Video_id }}) ( {{ .Updated_date }} 公開)
{{ end }}

