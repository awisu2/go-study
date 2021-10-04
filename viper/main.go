package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)


type ConfigB struct {
	B2 int `yaml:"b2"`

}

type Config struct {
	A string `yaml:"a"`
	B *ConfigB `yaml:"b"`
	
}

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	// ファイルを検索するパス
	// viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath(".") // current directory

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	a := viper.GetString("a")
	log.Println(a)

	// structに読み込む
	var config Config
	viper.Unmarshal(&config)
	log.Println(config, config.B) // {A 0xc000024a88} &{21}
}