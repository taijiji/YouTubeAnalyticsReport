# YouTubeAnalyticsReport
Add your [YouTube API key](https://developers.google.com/youtube/registering_an_application) and your YouTube Channel ID into `.env` file.

```
vi .env

API_KEY="XXXXXXXX"
CHANNEL_ID="XXXXXXXX"
```

Run `main.go`. This script will get the most recent 15 videos on your YouTube Channel.

```
$ go run main.go
==================================================
Channel Title:  show int インターネットの裏側解説
Channel ID:  UCpO3RcIrPaDJJ0Q3cbZrvZA
Suscribers: 3890
==================================================
----------------------------------------------------------------------------------------------------
Video Title:  日本中のネットワークエンジニアが集まる JANOG51 の会場から現場の雰囲気をレポートします
Video ID:  DS_XckbvoeM
Uploaded Date:  2023-02-06
View Counts:  524
Like Counts:  19
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  JANOG51終了直後！富士山からJANOG51現地の雰囲気をふりかえりました
Video ID:  7nwO7U6Ovlg
Uploaded Date:  2023-01-30
View Counts:  255
Like Counts:  12
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  ネットワークエンジニアが注目しておくべき最新技術カンファレンス 4選
Video ID:  FsUT4wX1L8A
Uploaded Date:  2023-01-23
View Counts:  373
Like Counts:  10
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  show int の2022年振り返り &amp; 2023年に挑戦したいこと
Video ID:  vn_1gqEh8hw
Uploaded Date:  2023-01-16
View Counts:  183
Like Counts:  7
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  【開催直前】JANOGスタッフがおすすめするJANOG51 Meeting みどころ紹介【最新技術から現地観光情報まで】
Video ID:  NFNNBOflnkQ
Uploaded Date:  2023-01-09
View Counts:  594
Like Counts:  15
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  Twitter・Facebookなどのグローバル企業で人員削減が相次いだ2022年。「大規模人員削減時代」にエンジニアがいま考えておくべきこと
Video ID:  MDqkONWiu1k
Uploaded Date:  2022-12-26
View Counts:  543
Like Counts:  10
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  ワールドカップ2022 の裏側でネットワークエンジニア達がヒヤヒヤしていた事情を解説します【ABEMAおよび各社エンジニアのみなさまお疲れさまでした】
Video ID:  iQg49y7KSVU
Uploaded Date:  2022-12-19
View Counts:  516
Like Counts:  14
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  トップレベルのエンジニアに求められるコミュニケーション力・リーダーシップ力
Video ID:  FkAUTYip1n8
Uploaded Date:  2022-12-13
View Counts:  301
Like Counts:  8
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  アジア各国を飛び回るスターエンジニア maz 松崎 吉伸さん【エンジニア対談】
Video ID:  YkvU8JmF2qw
Uploaded Date:  2022-12-05
View Counts:  433
Like Counts:  15
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  「ネットワークエンジニア未経験者歓迎！」は本当に未経験でもOKなのか
Video ID:  f4Rb4HIuBJI
Uploaded Date:  2022-11-28
View Counts:  461
Like Counts:  12
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  複数の会社から転職オファーをもらったときの選び方【視聴者さんからのお悩み相談】
Video ID:  hB4CBLmK0lw
Uploaded Date:  2022-11-21
View Counts:  233
Like Counts:  0
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  衛星インターネット Starlink が爆速な理由 を現役ネットワークエンジニアが解説してみた
Video ID:  Kfnd_iA_rm8
Uploaded Date:  2022-11-14
View Counts:  3792
Like Counts:  74
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  現役ネットワークエンジニアが実践する「モチベーションに頼らない成長戦略」
Video ID:  SGJVDqkquT4
Uploaded Date:  2022-11-07
View Counts:  339
Like Counts:  12
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  日本人が陥りがちな外国人エンジニアとのコミュニケーション失敗例
Video ID:  fHZpgFVRlNY
Uploaded Date:  2022-10-31
View Counts:  247
Like Counts:  9
Dislike Counts:  0
----------------------------------------------------------------------------------------------------
Video Title:  StackStorm で実現するネットワーク完全自動化の世界
Video ID:  Z8e5vso5SRU
Uploaded Date:  2022-10-24
View Counts:  486
Like Counts:  20
Dislike Counts:  0
```