# agouti

golang でのブラウザコントロール API サポート

- public: [Agouti](https://agouti.org/)
- github: [sclevine/agouti: A WebDriver client and acceptance testing library for Go](https://github.com/sclevine/agouti)
- test(sample): [agouti/internal/integration at master · sclevine/agouti](https://github.com/sclevine/agouti/tree/master/internal/integration)
- api: [agouti package \- github\.com/sclevine/agouti \- pkg\.go\.dev](https://pkg.go.dev/github.com/sclevine/agouti)

## メモ

- 特定のブラウザドライバーをコントロールするため、事前にドライバのインストールが必要
  - chrome: [ChromeDriver \- WebDriver for Chrome \- Downloads](https://chromedriver.chromium.org/downloads)
  - *failed to navigate: request unsuccessful: invalid session id*
    - ドライバのバージョンが違うと発生する。Driverの生成時ではなくNavigate()を実行したときにNavigateの失敗としてエラーがでる

## toto

- [] ginko, gomegaとの連携が可能

## オプション

driverには起動時にオプションを指定できる。指定パラメータは各ドライバに準じる

### chrome

- `--headless`, `--disable-gpu`: ヘッドレスモードで起動(max/linuxのみ)
  - [ヘッドレス Chrome ことはじめ  \|  Web  \|  Google Developers](https://developers.google.com/web/updates/2017/04/headless-chrome#cli)
- `--incognito`: シークレットモード
- `--no-startup-window`: タスクトレイ内で起動
- `--user-agent`: ユーザエージェントの変更
- `--window-size={width},{height}`: 起動時のwindowサイズ
- `--user-data-dir={dir}`: ユーザデータディレクトリ(cookie、ブラウザ設定、ログイン情報などブラウザが管理している情報を継続できる)

#### links

- [起動オプション \- Google Chrome まとめWiki](http://chrome.half-moon.org/43.html)
- こちらのほうが補完率が高い: [List of Chromium Command Line Switches « Peter Beverloo](https://peter.sh/experiments/chromium-command-line-switches/)
