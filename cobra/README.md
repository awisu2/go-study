# cobra

study cobra, golang cli module

- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
- [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)

## 簡単解説

- "cmd" パッケージを作成し、main から Execute()などで実行する構成
- 以下のようにコマンドに相当するインスタンスを生成し。Execute()することで Run が実行される
- コマンドにはコマンドを追加でき、階層的呼び出しが可能
  - `rootCmd.AddCommand(secondCmd)`
  - 一番親となるコマンドのことを rootCommand と称するらしい
- Execute()時に 同 package 内の init() が先行して呼び出される
  - ファイル分類してコマンドを分割している場合などは、ここで AddCommand することでファイル単位の独自性を確保できる

```go
var rootCmd = &cobra.Command{
  Use: "cobra",
  Short: "cobra sample",
  Logn: "try cobra sample",
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("hello")
  }
}
```

- Run は RunE にすることで error を返却することができる
- サンプルでサラッと出てくる viper は同じ人? が作った設定ファイルを読み込む module
- go の flag もサポート
  - `flag.Parse()` すると 引数が flag.Args() で取得できる
  - 事前に型宣言した値を作成しておくことによって、型付き変数に Parse させることも可能
