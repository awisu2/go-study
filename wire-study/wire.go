//go:build wireinject

package main

import "github.com/google/wire"

// 引数は型があっていれば良いみたい
// NewMessage need `message string`
func InitializeEvent(msg string) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}

// 条件で分けて異なるProviderを使用
func InitializeEventWithIf(msg string) Event {
	if msg == "" {
		wire.Build(NewEvent, NewGreeter, NewMessage2)
		return Event{}
	} else {
		wire.Build(NewEvent, NewGreeter, NewMessage)
		return Event{}
	}
}

func InitializeFooBar() FooBar {
	wire.Build(
		wire.NewSet(
			ProvideFoo,
			ProvideBar,
			wire.Struct(new(FooBar), "MyFoo", "MyBar"),
		),
	)
	return FooBar{}
}

func InitializeVoice() ManVoice {
	wire.Build(
		wire.Value("woooo"),
		ProvideNormalMan,
		wire.Bind(new(InterfaceMan), new(*NormalMan)),
		ProvideManVoice,
	)
	return ManVoice{}
}

var ManvoiceSet = wire.NewSet(
	wire.Value("woooo"),
	ProvideNormalMan,
	wire.Bind(new(InterfaceMan), new(*NormalMan)),
	ProvideManVoice,
)

func InitializeVoiceWithSet() ManVoice {
	wire.Build(ManvoiceSet)
	return ManVoice{}
}
