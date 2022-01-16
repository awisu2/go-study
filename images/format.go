package images

// formtディレクトリなど作ってサブパッケージに分けたいがimportが複雑になるので同パッケージで管理
type Format string

const (
	Jpg Format = "jpg"
	Png Format = "png"
)
