# sharedGo

go を.so ファイルにビルドし、他言語に組み込む勉強

1. 一定の要件を満たしたコードを記述
2. .so ファイルのビルド
   - `go build -o awesome.so -buildmode=c-shared .`
3. 各種言語から呼び出し

## links

- [Calling Go Functions from Other Languages | by Vladimir Vivien | Learning the Go Programming Language | Medium](https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf)
