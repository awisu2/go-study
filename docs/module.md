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
