# cobra

study cobra, golang cli module

- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
- [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)

## 簡単解説

- "cmd" パッケージ(ディレクトリ)を作成し、main から Execute()などで実行する構成
- コマンドに相当するインスタンスを生成し。Execute()する
- コマンドにはコマンドを追加でき、階層的呼び出しが可能
  - `rootCmd.AddCommand(secondCmd)`
  - 一番親となるコマンドのことを rootCommand と称するらしい
  - 親コマンドへの登録などは 各ファイルの init() で実行することでファイルごとに処理系を分散

## sample

```go
var rootCmd = &cobra.Command{
  Use: "cobra",
  Short: "cobra sample",
  Logn: "try cobra sample",
  // positional argumentsの設定(default: cobra.NoArgs(引数あるとエラー))
  Args: cobra.ArbitraryArgs,
  // これがない場合はヘルプが実行される
  Run: func(cmd *cobra.Command, args []string) {
    log.Println("hello")
    // フラグから特定の型で値を取得(Flags と PersistentFlags は明確に異なるので注意)
    foo, err := cmd.Flags().GetString("foo")
    bar, err := cmd.PersistentFlags().GetString("bar")
  }
}

func init() {
	cobra.OnInitialize(initConfig)

// 1. --author or -a の引数設定
	// 値取得は `cmd.Flags().GetString("author")`
	rootCmd.Flags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// 2. --config の引数設定、Persistent がつき下位のコマンドにも継続適用される
	// 値取得その1は `&cfgFile` にセットされている。見たまんま
	// 値取得その2は `cmd.PersistentFlags().GetString("config")` でも取得可能
	// Note: Flags と PersistentFlags は明確に別れているので注意
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// 必須にする
	rootCmd.MarkPersistentFlagRequired("config")

	// 子コマンド
  rootCmd.AddCommand(subCmd)
}
```

## 用語

- flag arguments: "-f", "--foo" などの -付きの引数
  - cobraでは主にこれの設定を行う
- positional arguments: flag argumentsでない直接(?)並べられた引数
- Persistent Flags: 下位のコマンドに継承されるフラグ引数
- Local Flags:

## もうちょっと細かい話

- positional argumentsを許容したい
  - Argsをコマンドに設定しないと、"unknown command" となりエラーになる
    - 後述する NoArgs がデフォルトで設定されているのだと思われる
  - Argsをサポートする関数郡(別に細かいことするつもりな石というときに変わりにセット)
    - 引数なし: cobra.NoArgs
    - とりあえずOK: cobra.ArbitraryArgs
    - 特定の引数のみOK: cobra.OnlyValidArgs
      - 別途ValidArgsを設定しておく必要あり、ValidArgsFunctionで更に細かい設定も可能
    - 最大/最小個数: cobra.MinimumNArgs(int), cobra.MaximumNArgs(int), cobra.RangeArgs(min, max)
    - ピッタリ個数: cobra.ExactArgs(int), cobra.ExactValidArgs(int)
- Run は RunE にすることで error を返却することができる
- サンプルでサラッと出てくる viper は同じ人? が作った設定ファイルを読み込む module
- go の flag もサポートというか、flagてき挙動がメイン
  - `flag.Parse()` すると 引数が flag.Args() で取得できる
  - 事前に型宣言した値を作成しておくことによって、型付き変数に Parse させることも可能
