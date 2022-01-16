# image

go のライブラリ [image package \- image \- pkg\.go\.dev](https://pkg.go.dev/image) の勉强。
リサイズパッケージは長期間メンテ無しとのこと。[nfnt/resize: Pure golang image resizing](https://github.com/nfnt/resize)

- 基本型は: `image.Image`
  - 生成の際は `image.NewRGBA() *image.RGBA` や `image.NewCMYK() *image.CMYK` などで拡張型が取得されることもしばしば
- ファイル保存: Fileストリーム に対して `image/jpeg` や `image/png` など指定の形式で `Encode(w io.Write, m image.Image)` で画像オブジェクトから書き込む
- リサイズ: `golang.org/x/image/draw` により処理可能(後述)
- 色塗り: Set() 関数にて 1ピクセルずつ塗るのが基本。色指定は color.Color インスタンス
  - 画像オブジェクトのように `color.RGBA` や　`color.CMYK` などの拡張型が用意されており、そちらから生成するのが基本か

## links

- [image package \- image \- pkg\.go\.dev](https://pkg.go.dev/image)
- [draw package \- golang\.org/x/image/draw \- pkg\.go\.dev](https://pkg.go.dev/golang.org/x/image/draw)
- [color package \- image/color \- pkg\.go\.dev](https://pkg.go.dev/image/color)

## draw

画像オブジェクトに対して書き込み処理を行える

- リサイズ
  - 画像オブジェクトに対して書き込むので、先に目的サイズのオブジェクトを用意する必要がある
  - 行うのは、もと画像をscaleするという処理
- Kernel: カーネルは、対称カーネル関数によって重み付けされたソースピクセルをブレンドする補間器です。(翻訳まま)
  - インスタンスを生成する際にパラメータを指定する必要がある
    - 事前用意されたインスタンスに `BiLinear`, `CatmullRom` があるのでこれらを使うとよい
      - 両方とも遅いが大抵高クオリティの結果を出すとのこと
      - [draw package \- golang\.org/x/image/draw \- pkg\.go\.dev](https://pkg.go.dev/golang.org/x/image/draw#Kernel)
    - todo: 詳細不明(コンピュータグラフィックの領域)
  - `Scale()`: 特定の画像オブジェクトをサイズ変更して書き込み
  - `NewScaler()`: Scale() を持つ、Scallerインスタンスを生成
    - パラメータをいくつか渡すが、Scale() 関数の引数が減らないので用途不明。"A Scaler is safe to use concurrently." とあるので goroutine用？
  - `Copy()`: Scale() に似ているが、画像の拡縮なしにdstの指定ポイントを起点として書き込む
  - `DrawMask()`: todo: 調査不足
  - `Draw()`: todo: 調査不足 DrawMask() を mask = nil で実行

### Scale() の挙動

- 色が先に塗られていた場合dstが優先される
- 第5引数の draw.Op: alphaが効く形式のときに有効、指定できるのは下記2つ
  - draw.Over: 上書き(色は交じる)
  - draw.Src: dstが優先され色の合成はされない

## color

- color.RGBA
  - [color package \- image/color \- pkg\.go\.dev](https://pkg.go.dev/image/color#RGBA.RGBA)
  - 全パラメータ uint8
    - 0-255 で扱われ、Aは 0 の時透明
