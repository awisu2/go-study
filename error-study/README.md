# error

- [errors package \- errors \- pkg\.go\.dev](https://pkg.go.dev/errors)
- [A Tour of Go](https://go-tour-jp.appspot.com/methods/20)

```go
// basic interface
type error interface {
    Error() string
}

// for errors interface(recomend)
type CustomError interface {
    Error() string
    Is(target error) bool
    Unwrap() error
}

// fmt can create error with `Unwrap() error` function
err := errors.New("base error")
errWrap := fmt.Errorf("wrap = %w = wrap.", err)
fmt.Println(errWrap) // wrap = base error = wrap.

// check error (Is, As)
var cErr *CustomError
if errors.As(err, &cErr) {
  fmt.Printf("has same type error. detail %v\n", cErr)
}

var cErr CustomError{}
if errors.Is(err, cErr) {
  fmt.Printf("has same error. == %v\n", err)
}
```

## NOTE

- 通常の error は error() string を持つ ingerface
- error のチェックは != nil で行う
- ハンドリングは `error`の type で判断する
- 各 `type` に `Is()`を実装することで、`erros.Is()` が利用できるようになる
- 各 `type` に `Unwrap()` メソッドを実装することで、`errors.As()`, `errors.Is()` wrap しているエラーも対象となる
  - `func (err anError) Unwrap() error` 関数で返却可能な場合 wrap しているという
- `errors.As() bool`: 小階層に渡って、型が同じエラーが存在するかをチェック
- `errors.Is() bool`: 小階層に渡って、`Is()`メソッド(`==`)によって一致するエラーが存在するかをチェック
