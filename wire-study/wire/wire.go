//go:build wireinject

package wire

import (
	ws "github.com/awisu2/go-study/wire-study"
	"github.com/google/wire"
)

func InitializeEvent(msg string) ws.Event {
	wire.Build(ws.NewEvent, ws.NewGreeter, ws.NewMessage)
	return ws.Event{}
}
