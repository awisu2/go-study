# sharedGo

go を.so ファイルにビルドし、他言語に組み込む勉強

1. 一定の要件を満たしたコードを記述
2. バイナリビルド (.so or .dll, .h の２ファイルが生成される)
   - `go build -buildmode=c-shared -o awesome.so .`
   - NOTE: node の場合は dll ファイルでないと エラーになる "Error: Dynamic Linking Error: Win32 error 126"
     - `go build -buildmode=c-shared -o awesome.dll .`
     - なぜかはわからないが、node ＋ windows だからかな？
     - c の場合は.so で大丈夫だった
3. 各種言語から呼び出し(fromXXX となったディレクトリで試している。)
   - リポジトリに乗らないので、上記ビルドしてコピーのこと

## links

- [Calling Go Functions from Other Languages | by Vladimir Vivien | Learning the Go Programming Language | Medium](https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf)
