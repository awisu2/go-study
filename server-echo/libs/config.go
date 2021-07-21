package libs

import (
	"go-study/server-echo/models"
)

var DbConfigs = models.Configs{
	"defalt": {
		Key: "default",
		Kind: models.DB_KIND_SQLITE,
	},
}