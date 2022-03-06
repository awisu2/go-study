# 基本事項

用語やかんたんなコマンド事前に知っておきたいこと

- go 自体の install: [Download and install - The Go Programming Language](https://golang.org/doc/install)
- commands: 詳細は後述
  - create mod: `go mod init {domain/module}`
  - update packages: `go mod tidy`
  - run mod: `go mod .`
- 環境変数: 後述
- [test](../test)
- [getting start/hello world](../helloworld)
  - 公式の勉强環境: [A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)

## install

### linux

```bash
# install
GOFILE="go1.16.5.linux-amd64.tar.gz"
cd
wget https://golang.org/dl/${GOFILE}
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf ${GOFILE}

# add path (if not exists)
cat - << 'EOF' > ~/.bash_profile
# go
export PATH=$PATH:/usr/local/go/bin
EOF

# after restart
go version
```

## 環境変数

- GOPATH: go が path 解決する際に root として扱うディレクトリ(ワークスペース)
  - 配下に package やバイナリ、開発コードが存在することが前提として動作する
  - 利用可能なディレクトリであればどこでも OK
  - os(process)ごとに１つしか指定できない
    - (開発元の異なるなど)ディレクトリを分けたい場合逐一切り替える必要あり
- GOROOT: よくわからないが、go 本体のインストールパスが割り当てられていた
  - GOROOT の設定は不要との記事は散見する。実際しなくて問題ない
- GOBIN: built 時の install 先 (指定しなければ `${GOPATH}/bin` になるっぽい)

## 困り毎

- vscode での補完がおかしい: vscode workspace root = go pacakgae root である必要があるとのこと
  - このリポジトリの用に小ディレクトリに package 置いていると発生する
  - add folder to workspace で root になるようにしてやれば OK

## 用語

- module と package の違い
  - package: 同ディレクトリ内の go ファイル群。go ファイルの先頭での package 宣言名が同じもの
  - module: go.mod によって管理される、package 群。install/get の対象
