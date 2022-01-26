# module について

- [module と package の違い](./basics.md)
- `go.mod`ファイルによって管理される
- コマンド
  - `go mod init [module]`: go.mod ファイルの生成 module に指定した名前がセットされる
  - `go install {module}`: バイナリのビルドとインストール
    - 提供されている既存モジュールコマンドを実行したい時用
  - `go get {module}`: 対象モジュールのダウンロードと go.mod の更新 (開発時は install ではなくこちら)
  - `go mod tidy`
