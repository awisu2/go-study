package libs

import (
	"go-study/server-echo/dbs"
)

var DbConfigs = dbs.Configs{
	"defalt": {
		Key: "default",
		Kind: dbs.DB_KIND_SQLITE,
	},
}