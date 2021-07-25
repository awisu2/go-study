package app

import (
	"go-study/echo-middleware/libs/echo"
)

type Config struct {
	Title string
	EchoConfig *echo.Config
}

var _config *Config

// cycle importを極力避けるため、type宣言とインスタンスの生成を分離
//
// 目的：ここをsingletonのinstance倉庫にしたいため
// 課題：viewでアプリ全体のconfigインスタンスを必要としておりその際にrouterと衝突してcycle importになる
// 必要な箇所に、インスタンスをセットすることである程度は防げるが、今後どこかで必要になるたびにセットするのもなにか違う気がするのでこの対応
// 
func SetConfig(config *Config) {
	_config = config
}

func GetConfig() *Config {
	return _config
}
