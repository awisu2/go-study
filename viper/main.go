package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ContentDir string
	B          string
}

func main() {
	sampleReadEnv()
}

func sampleSetDefault(def string) string {
	// set default
	viper.SetDefault("mykey", def)

	// get value
	return viper.GetString("mykey")
}

func sampleReadFile() Config {
	/* read file
	 * ConfigPathは複数指定できるが、読み込む(管理する)ファイルは1つだけ
	 */
	viper.SetConfigName("config.json") // file name
	viper.SetConfigType("json")        // ファイル名に拡張子がないときに必要
	viper.AddConfigPath("./")          // target directory
	err := viper.ReadInConfig()
	fmt.Println("----- a")
	if err != nil {
		panic(fmt.Errorf("Fatal error config file. %w \n", err))
	}

	// read to struct
	var config Config
	viper.Unmarshal(&config)

	return config
}

func sampleReadEnv() string {
	/* read from environment
	 * 実行サンプル: `B=abc go run .`
	 */
	viper.AutomaticEnv()
	return viper.GetString("B")
}
