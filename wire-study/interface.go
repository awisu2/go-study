package main

type InterfaceMan interface {
	Voice() string
}

type NormalMan struct {
	Word string
}

func (m *NormalMan) Voice() string {
	return m.Word
}

type ManVoice struct {
	Voice string
}

func ProvideNormalMan(word string) *NormalMan {
	return &NormalMan{
		Word: word,
	}
}

func ProvideManVoice(man InterfaceMan) ManVoice {
	return ManVoice{
		Voice: man.Voice(),
	}
}
