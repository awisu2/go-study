# wire study

- [google/wire: Compile\-time Dependency Injection for Go](https://github.com/google/wire)
  - [tutorial](https://github.com/google/wire/blob/main/_tutorial/README.md)

```bash
# install
go get github.com/google/wire/cmd/wire

# generate wire_gen.go
# `wire` == `wire gen .`
wire
# spesific package
# `wire ./wire` == `wire gen ./wire`
wire ./wire

# regenerate (Only works if there is wire_gen.go)
go generate

# run
go run .
```

## NOTE

- why use wire: I think useful to use architectures such as DDD and clean architecture.
- always generate **wire_gen.go**
- not foreget `// +build !wireinject` or `//go:build wireinject`
- commands
  - help: `wire help` help
  - diff: `wire diff` output a diff between existing wire_gen.go files and what gen would generate.
  - show: `wire show` describe all top-level provider set.
- behavior
  - arguments: Determinded by type, not argument name.
    - [Binding Interfaces](https://github.com/google/wire/blob/main/docs/guide.md#binding-interfaces)
  - Can generate from multi files. However, files that do not use `wire` are not included. (from my try)
- `wire.Build()`
  - `wire.NewSet`は `wire.Build` で利用可能なSetを作成。事前に一定のパターンをまとめて置くことができる
  - `Wire.Bind` は interface を特定の 型と紐付ける。後続のproviderがinterfaceを必要とする場合に使用
  - `wire.Value` は 特定の型をBuild内でセットすることができる。(ex: provide string `wire.Value("str")`)

## problem

- Q: how to just set value in `wire.Build()`?
  - ex: want only set route to webframework engine
  - A: Even if don't use it, return one of the type value.
