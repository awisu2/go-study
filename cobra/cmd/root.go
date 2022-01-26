package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 受け取る引数の事前宣言
//
// - `cmd.PersistentFlags().GetString("{key}")`で取得も可能
// - 複数コマンド間での重複を考えるとstructのほうがいいかも
//
var (
	// Used for flags.
	cfgFile     string
	userLicense string
)

// コマンドの作成
//
// コマンドにコマンドを追加できるが、起点をrootCmdというのがデファクトらしい
//
var rootCmd = &cobra.Command{
	Use:   "sample", // コマンド名(ここでは "{go run .} sample" となる)
	Short: "short description",
	Long:  `long description`,

	// 引数に対する動作設定
	//
	// エラー返却時Runは実行されない
	//
	// 選択肢
	// - cobra.OnlyValidArgs: 指定した引数以外があるとエラー
	// - cobra.NoArgs(default): 引数があるとエラー
	// - cobra.ArbitraryArgs: エラーは返さない (Arbitrary = 任意)
	// - カスタム: https://pkg.go.dev/github.com/spf13/cobra#PositionalArgs
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	log.Println("args", args)
	// 	return nil
	// },
	Args: cobra.OnlyValidArgs,

	// 実行コマンド
	Run: func(cmd *cobra.Command, args []string) {
		// 引数の取得(flag式)
		author, _ := cmd.Flags().GetString("author")
		a, _ := cmd.Flags().GetString("a")
		config, _ := cmd.PersistentFlags().GetString("config")

		log.Printf("author: %v, a: %v, config: %v\n", author, a, config)

		log.Println("hello", cfgFile, userLicense, cmd.PersistentFlags().Lookup("author"), args)
	},
}

// initの実行順序は、コンパイラに渡された順序とのこと
// [go \- Init order within a package \- Stack Overflow](https://stackoverflow.com/questions/32829538/init-order-within-a-package)
// [The Go Programming Language Specification \- The Go Programming Language](https://go.dev/ref/spec#Package_initialization)
func init() {
	// 引数の処理後(cobra.Command.Runの前)に実行されるイベント
	cobra.OnInitialize(initConfig)

	// 引数設定
	//
	// 簡易説明
	//
	// - Flags: このコマンドのみのフラグ(引数)指定
	// - Persistent(=持続的)Flags: 配下(子)コマンドにも引数が引き継がれる(Persistent Flags, Local Flagsと呼ばれる)
	//
	// Flags/PersistentFlags:
	// - xxxVar: ショートハンド名の設定が可能
	// - xxxP: それぞれの型にあった(参照)引数へ値がセットされる
	// - xxxVarP: Var と P の併用
	// - xxx: Var、Pのどちらでもない最小限設定
	// - 引数名は "" にすることで、引数指定を無しにできる
	//
	// 必須にする: キーを引数宣言とは別に必須設定する `cmd.Mark[Persistent]FlagRequired("name")`
	//

	// 1. --author or -a の引数設定
	// 値取方法: `cmd.Flags().GetString("author")` ショートハンド名("a")では取得できない
	rootCmd.Flags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")

	// 2. 配下に引き継がれる引数指定(Persistent)、引数へのセット式(xxxVar)、+ 必須指定(Requiredを付与)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// viperによりコンフィグ設定からの値をセット
	//
	// viperはconfigファイルや環境変数を取り込むモジュール
	// Lookupで返却されるのは *pflag.Flag
	//
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	// 小コマンドの追加
	rootCmd.AddCommand(tryCmd)
}

func initConfig() {
	fromViper()
}

// cfgFileが設定されていた場合、viperにより値を読み込む
//
// TOOD: 調査不足
//
func fromViper() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	a := viper.GetString("a")
	log.Println(a)

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() error {
	return rootCmd.Execute()
}

var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := someFunc(); err != nil {
			return err
		}
		return nil
	},
}

func someFunc() error {
	log.Println("i'm some!")
	return nil
}
