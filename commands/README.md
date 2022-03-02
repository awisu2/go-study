# commands

```bash
go mod {module}
go mod tidy
go run .

go install {module}
```

go には module モードと、GOPATH モードがある。デフォルトでは module モードで、GOPATH モードは GOPATH 配下にすべてを詰め込む旧仕様。
module モードにすることで、module のバージョンなどを分離できる。。。らしい。 go.mod がその役目？

- mod
    - `go mod init {domain/module}`
    - `go mod tidy`
    - after writed `imoprt` in code
    - `go mod edit {command}`
    - replace module path: `go mod edit -replace example.com/greetings=../greetings`
- build and run
    - `go build`: build to executable file
    - ex: `go build main.go && ./main` (create **main** or **main.exe**)
    - ex2: `go build .` build by module(directory). (this sample create **commands** or **commands.exe**)
    - check install path (it builded): `go list -f {{.Target}}`
        - これを環境変数の PATH に入れておけばビルドしたコマンドが利用できる
    - `go run`: run with build. simply way at develop
    - ex: `go run .`

## when we use go build?

I found the "go build" usage in gin, it is web framework.
there can replace encoding/json (default) package.

[gin\-gonic/gin. Build with json replacement](https://github.com/gin-gonic/gin#build-with-json-replacement)

```bash
#jsoniter
go build -tags=jsoniter .
# go-json
go build -tags=go_json .
```
