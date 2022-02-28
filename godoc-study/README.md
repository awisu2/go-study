# go doc

- [Godoc: documenting Go code \- The Go Programming Language](https://go.dev/blog/godoc)
- [doc command \- cmd/doc \- pkg\.go\.dev](https://pkg.go.dev/cmd/doc)
- [godoc command \- golang\.org/x/tools/cmd/godoc \- pkg\.go\.dev](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

**NOTE**

- on cli, `doc` is now available. If use it, `go doc`.
  - `godoc` has been removed. ([x/tools/cmd/godoc: remove CLI support in Go 1\.12 · Issue \#25443 · golang/go](https://github.com/golang/go/issues/25443))
- [] can't generate html from doc command?

## sample

```bash
# print document all
go doc -all

# print document with search
go doc -c Hello
go doc -cmd Hello

# print document just symbol
go doc -cmd
```

## output of `go doc -all`

```txt
FUNCTIONS

func Hello(name string) string
    Hello print hello with name

func Hi(name string) string
    Hi print hi with name

    name string


TYPES

type Human struct {
        Name string // name
        // age
        // from 0 to 200
        Age string
}
    Human has human's variable
```

### install godoc command

```bash
go install golang.org/x/tools/cmd/godoc
no required module provides package golang.org/x/tools/cmd/godoc; to add it:
        go get golang.org/x/tools/cmd/godoc
```

but not exists

solution

```bash
go get golang.org/x/tools/cmd/godoc
```
