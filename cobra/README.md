# cobra

study cobra, golang cli module

- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
- [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper)

## 残タスク

- [] viperとの連携の細かい解説(必要になったら)
- [] なぜinit()で記載しているのか？

## 簡単解説

- pythonなどから来た場合、argparseでもいいかもしれない
  - [akamensky/argparse: Argparse for golang. Just because `flag` sucks](https://github.com/akamensky/argparse)
  - 優位点としては、viperとも連携できること。packageへの分割が前提として考慮されており見通しが良いことなどが挙げられる
- viper(config系module), flag と共存することも可能
- RunEを利用すべきか？: エラー時の処理を決めていない場合は、Runで良いと思われる
  - 個々のコマンドでエラー条件が変わるため、エラーを返してどうするという問題が残る
- viperとの連携: 細かいことは省くが、viperに値を寄せていく感じ。
 - [spf13/viper: Go configuration with fangs](https://github.com/spf13/viper#working-with-flags)
 - [cobra/user\_guide\.md at master · spf13/cobra](https://github.com/spf13/cobra/blob/master/user_guide.md#create-rootcmd)

### 参考コード

コピペ利用に対応

```go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var params = struct {
	arg1 int
}{}

// create command
var Cmd = &cobra.Command{
	Use:   "sample",
	Short: "short description",
	Args:  cobra.ArbitraryArgs, // arguments setting.(ArbitraryArgs: any value ok)
	// can chooses `Run()` or `RunE()`
	// args has cli parameters without has -{-str} option.
	//
	// get arg: cmd.Flags().GetInt("num")
	// get persistent arg: cmd.PersistentFlags().GetString("str")
	//
	// show help: `cmd.Help()`
	Run: func(cmd *cobra.Command, args []string) {
		_arg1, _ := cmd.Flags().GetInt("num")
		_arg2, _ := cmd.PersistentFlags().GetString("str")
		fmt.Printf("arg1(var): %v, arg1: %v, arg2: %v, args: %v\n", params.arg1, _arg1, _arg2, args)
		// cmd.Help()
	},
}

// run initialize like viper
func initConfig() {
}

// execute at running.
func init() {
	// run between parse arguments and run Command.Run()
	cobra.OnInitialize(initConfig)

	// arguments
	flags := Cmd.Flags()
	flags.IntVarP(&params.arg1, "num", "n", 99, "arg1 number")

	// persistent arguments
	pFlags := Cmd.PersistentFlags()
	pFlags.StringP("str", "s", "abc", "arg2 string")

	// require settings
	requireds := []string{"output", "width", "height"}
	persistentRequireds := []string{}
	for _, required := range requireds {
		Cmd.MarkFlagRequired(required)
	}
	for _, required := range persistentRequireds {
		Cmd.MarkPersistentFlagRequired(required)
	}

	// add command
	Cmd.AddCommand(subCmd)
}

// Execute(): Tipycaly function name
func Execute() error {
	return Cmd.Execute()
}

// create other command.
// (recomend) create on other package.
var subCmd = &cobra.Command{
	Use: "sub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}
```


## 実装

- サンプルコード: [cmd/root.go](./cmd/root.go)
- 配置: "cmd" パッケージ(ディレクトリ)内に記載していくのが参考構成
  - `init()` で引数設定、配下コマンドの追加をする理由: よくわかってない 引数、viperでのconfig値による処理の分岐を許容数ため？
- cmd.Execute() での統一を狙って、cmd 直下に `func Execute()` を用意しておく
- 基本は `&cobra.Command{}` でコマンドを作成し引数やdescriptionの設定、配下にコマンドを追加も可能
  - コマンドの追加: `rootCmd.AddCommand(secondCmd)`
- 引数の取得: `cmd.[Persistent]Flags().GetXXX({key})`で取得が可能
  - 引数設定で直接引数にセットすることも可能

## cobra.Commandのもうちょっと細かい話

公式: [cobra package \- github\.com/spf13/cobra \- pkg\.go\.dev](https://pkg.go.dev/github.com/spf13/cobra#Command)

- Run[E]とUseの省略: コマンド無しでhelp、第一コマンド名をサブコマンドのUseにできる(rootCmdにおすすめ)
- Args: 引数の挙動を指定。エラーを返却するとRunが実行されない
  - cobra.OnlyValidArgs: 指定した引数以外があるとエラー
    - 別途ValidArgsを設定しておく必要あり、ValidArgsFunctionで更に細かい設定も可能
  - cobra.NoArgs(default): 引数があるとエラー
  - cobra.ArbitraryArgs: エラーは返さない (Arbitrary = 任意)
  - 最大/最小個数: cobra.MinimumNArgs(int), cobra.MaximumNArgs(int), cobra.RangeArgs(min, max)
  - ピッタリ個数: cobra.ExactArgs(int), cobra.ExactValidArgs(int)
  - カスタム: https://pkg.go.dev/github.com/spf13/cobra#PositionalArgs
    ```go
      Args: func(cmd *cobra.Command, args []string) error {
        log.Println("args", args)
        return nil
      },
    ```
- RunE/Run: RunEでerror返却可能。両方設定された場合RunEが優先
  - `[Persistent]{Per|Post}Run[E]()` も存在 (persistent = 持続的)
  - type: `func(cmd *Command, args []string) error`
    - argsには、第1要素にコマンド名, 第2要素以降に、positional argumentsがセットされる
  - 引数の取得: 例: `cmd.PersistentFlags().GetString("config")`
    - キー名はlong名のみで取得可能
    - また、別途引数と連携するように設定した変数を利用することも可能(後述)

### 引数、配下コマンド設定

- 公式: [pflag package \- github\.com/spf13/pflag \- pkg\.go\.dev](https://pkg.go.dev/github.com/spf13/pflag#FlagSet)

- int()の実行順序: コンパイラに渡された順序とのこと
  - 実質順序不明なので気にしないでいいように書くこと
  - 参考
    - [go \- Init order within a package \- Stack Overflow](https://stackoverflow.com/questions/32829538/init-order-within-a-package)
    - [The Go Programming Language Specification \- The Go Programming Language](https://go.dev/ref/spec#Package_initialization)
- 引数の設定:
  - 参考構文: `cmd.Flags().String("abc", "myAbc", "abc word")`
    - コマンド実行時 --abc で指定できる,  型はString, 初期値: "myAbc", help時説明: "abc word"
  - 構文ルール: `cmd.[Persistent]Flags().{type}[Var][P]($arg, "long", "short", "myAbc", "abc word")`
    - PesistentFlags: 配下コマンドでも有効
    - Varを付与: ショートハンド名の設定が可能
    - Pを付与す: 変数への値セットが可能
- 必須にする: `cmd.Mark[Persistent]FlagRequired("name")`
  - 引数設定とは別に必要

## 用語

後述する記載で突然出てきたときに

- flags: goではこのモジュールで引数を解析する
- flag arguments: "-f", "--foo" などの -付きの引数
  - cobraでは主にこれの設定を行う
- positional arguments: flag argumentsでない直接(?)並べられた引数
- Persistent: 下位のコマンドに継承されるフラグ引数
