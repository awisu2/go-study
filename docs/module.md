# module について

とりあえずわかったところから記載しているので間違っている可能性あり

- [module と package の違い](./basics.md)
- `go.mod` ファイルによって管理される
- コマンド
  - `go mod init [module]`: go.mod ファイルの生成 module に指定した名前がセットされる
  - `go mod tidy`: ディレクトリ配下の go ファイルで import されているモジュールを対象に、go.mod の更新及び、インストールが行われる。利用されていないものは go.mod から削除される
  - `go install {module}@{version}`: バージョン指定をして、モジュールをインストールできる
    - go.mod の更新は行わない
    - version: `latest`, `v1.0.0` の用に指定可能
    - 特定のバージョンをグローバルインストール(コマンド利用用途に思われる？)
  - `go get {module}[@{version}]`: 対象モジュールのインストールと go.mod の更新 (開発時は install ではなくこちら)を行う。
    - `go install` とほぼ同じ動作なのでグローバルインストールも行われる
    - version の指定なしの場合は _@latest_ を対象に実行される
    - 目的がモジュール内での利用のため、install の用に main package がないなどのエラーはでない
- `go get`, `go install` されたモジュールの行き先
  - GOPATH 配下にバージョン毎にディレクトリに保存されており、必要なときに再利用される
  - [] TODO: 最新化の方法

## private github から `go get` する

ポイントはこちら

1. github の httpsアクセス時に personal access tokensを噛ませて、privateでも通るようにする
   - go get は httpsアクセスをしているとのこと。まずはgithubの設定
2. GOPRIVATE 環境変数に対象のmodule名を登録
   - 自分のであれば, "github.com/awisu2/*" またはいつもgoを頭につけるので "github.com/awisu2/go*"
     - "," 区切りで複数設定可能、メタ文字が効くらしい

### 実践

personal access tokens を githubで発行

[個人アクセストークンを使用する - GitHub Docs](https://docs.github.com/ja/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)

gitに設定

```bash
TOKEN="{githubAccessToken}"
git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
```


f: ~/.gitconfig

以下の設定が追加される

```gitconfig
...
[url "https://{githubAccessToken}:x-oauth-basic@github.com/"]
        insteadOf = https://github.com/
```

ここまで、githubは通っているがgoのセキュリティでストップする。
環境変数 GOPRIVATE に許可するmoduleのパターンを設定

```bash
echo $GOPRIVATE
github.com/awisu2/*
```
