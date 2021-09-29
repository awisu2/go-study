package echo

import (
	"fmt"
	"html/template"
	"log"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type StaticConfig struct {
	Prefix string
	Root string
}

type Router func(e *echo.Echo)

type TemplateConfig struct {
	Key string
	// tempalteに適応する関数郡
	CreateFuncs func(t *template.Template) map[string]interface{}
	// template.ParseGlobによる対象テンプレートファイル(ex: "templates/*.html")
	GlobPattern string
}

// structのデフォルトがfalseのため、デフォルトonのものはNo設定
type Config struct {
	IsProduction bool
	Address string
	Router Router
	NoRecover bool
	NoLogger bool
	Static *StaticConfig
	IsCors bool
	CorsConfig *middleware.CORSConfig
	Template *TemplateConfig
	IsFixLocationForWindows bool // windowsの場合localhostをつける
}

func Create(config *Config) *echo.Echo {
	e := echo.New()

	if config.IsFixLocationForWindows {
		AddLocalhostIfWindows(config)
	}

	if !config.NoRecover {
		log.Println("use middleware Recover")
		e.Use(middleware.Recover())
	}

	if !config.NoLogger {
		log.Println("use middleware Logger")
		e.Use(middleware.Logger())
	}

	if config.Static != nil {
		log.Println("set static")
		e.Static(config.Static.Prefix, config.Static.Root)
	}

	if config.IsCors {
		log.Println("set cors")
		setCors(e, config)
	}

	if config.Template != nil{
		seTemplate(e, config)
	}

	if config.Router != nil {
		log.Println("run router")
		config.Router(e)
	}

	return e
}

func StartWithCreate(config *Config) *echo.Echo {
	e := Create(config)
	e.Logger.Fatal(e.Start(config.Address))
	return e
}


func setCors(e *echo.Echo, config *Config) {
	if config.CorsConfig == nil{
		e.Use(middleware.CORS())
	} else {
		e.Use(middleware.CORSWithConfig(*config.CorsConfig))
	}
}

// windowsの場合起動ごとにセキュリティアラートが出るための回避
func AddLocalhostIfWindows(config *Config) {
	if runtime.GOOS == "windows" {
		fmt.Println("add localhost")
		config.Address = "localhost" + config.Address
	}
}

func seTemplate(e *echo.Echo, config *Config) {
	t := template.New(config.Template.Key)

	if config.Template.CreateFuncs != nil {
		t = t.Funcs(config.Template.CreateFuncs(t))
	}

	e.Renderer = &Template{
		// template.Must: 解析時のエラーキャッチ、第2引数(err) != nil のときpnic
		// template.ParseGlob: パターンにより対象ファイルからテンプレートインスタンスを生成する
		templates: template.Must(t.ParseGlob(config.Template.GlobPattern)),
	}
}

