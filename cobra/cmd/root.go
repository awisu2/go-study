package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
)

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "A generator for Cobra based Applications",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// -f, --foo などフラグ引数以外の引受数をここで受け取れる
	// この設定がない場合フラグ引数以外があるとエラーになる
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	log.Println("args", args)
	// 	return nil
	// },
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		author, _ := cmd.Flags().GetString("author")
		log.Println(author)
		config, _ := cmd.PersistentFlags().GetString("config")
		log.Println(config)

		log.Println("hello", cfgFile, userLicense, cmd.PersistentFlags().Lookup("author"), args)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// NOTE: Persistentを付与すると、小コマンドでも引数設定が適用される(Persistent Flags, Local Flagsと呼ばれる)
	// NOTE: "" にすることで、引数指定を無しにできる
	// NOTE: xxxVar とするとこでそれぞれの型にあった参照引数へ値がセットされる
	// NOTE: がキー名として利用される
	// 必須にする: `cmd.MarkPersistentFlagRequired("name")` or `cmd.MarkFlagRequired("name")` を別途宣言

	// 1. --author or -a の引数設定
	// 値取得は `cmd.Flags().GetString("author")`
	rootCmd.Flags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// 2. --config の引数設定、Persistent がつき下位のコマンドにも継続適用される
	// 値取得その1は `&cfgFile` にセットされている。見たまんま
	// 値取得その2は `cmd.PersistentFlags().GetString("config")` でも取得可能
	// Note: Flags と PersistentFlags は明確に別れているので注意
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	// 対象を必須にしたい
	rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// viperはconfigファイルや環境変数を取り込むモジュール
	// Lookupで返却されるのは *pflag.Flag
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(tryCmd)
	// rootCmd.AddCommand(initCmd)
}

func initConfig() {
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
