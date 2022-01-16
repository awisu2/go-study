# image

go のライブラリ [image package \- image \- pkg\.go\.dev](https://pkg.go.dev/image) の勉强。
リサイズパッケージは長期間メンテ無しとのこと。[nfnt/resize: Pure golang image resizing](https://github.com/nfnt/resize)

- 基本型は: `image.Image`
  - 生成の際は `image.NewRGBA() *image.RGBA` や `image.NewCMYK() *image.CMYK` などで拡張型が取得されることもしばしば
- ファイル保存: Fileストリーム に対して `image/jpeg` や `image/png` など指定の形式で `Encode(w io.Write, m image.Image)` で画像オブジェクトから書き込む
- リサイズ: `golang.org/x/image/draw` により

## links

- [image package \- image \- pkg\.go\.dev](https://pkg.go.dev/image)
- [draw package \- golang\.org/x/image/draw \- pkg\.go\.dev](https://pkg.go.dev/golang.org/x/image/draw)
