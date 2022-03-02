# commands

```bash
go mod {module}
go mod tidy
go run .

go install {module}
```

- `go build`: build to executable file
  - ex: `go build main.go && ./main` (create **main** or **main.exe**)
  - ex2: `go build .` build by module(directory). (this sample create **commands** or **commands.exe**)
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
