# spotify-BPM-sorter

- Spotify Web API を使ってプレイリスト内の楽曲を BPM 順に並び替えて標準出力する
- Sort songs in a playlist by BPM using Spotify Web API and standard output

## Setup

- Spotify登録: Spotify for Developers でアプリを作成し、Client ID と Client Secret を取得
- 環境変数ファイル: リポジトリ直下に .env を作成し、次の2つを記入（`<>`記号は不要）
  - POTIFY_ID=<Client ID>
  - SPOTIFY_SECRET=<Client Secret>

## Run

- Goで実行: go run . <プレイリストID> を実行
  - 指定したプレイリストに含まれるトラックIDが標準出力へ
    - 現状はプレイリスト内トラックの Spotify トラックIDを上から順に表示
    - `main.go` 内の `fmt.Println(id)` を編集すると曲名や追加日時など別の情報も出力可能
- プレイリストID取得: 
  Spotify アプリ／Web の共有リンクから抜き出す
  - 例: `https://open.spotify.com/playlist/xxxxxxxx?si=...` から xxxxxxxx 部分を抜き出し

## Thanks

- [Dashboard \_ Spotify for Developers](https://developer.spotify.com/dashboard)
- [Web API \_ Spotify for Developers](https://developer.spotify.com/documentation/web-api)
- [Go で Spotify Web API を叩いてみる](https://zenn.dev/shimpo/articles/trying-spotify-api-with-go)
- [Web API Reference \_ Spotify for Developers](https://developer.spotify.com/documentation/web-api/reference/get-audio-features)
- [spotify package - github.com_zmb3_spotify_v2 - Go Packages](https://pkg.go.dev/github.com/zmb3/spotify/v2)
