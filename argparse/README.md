# argparse

- golang のデフォルトで、flag という機能が備わっている
- argparse を利用するメリット
  - flag より細かい設定が可能
  - os.Args 以外を Parse 設定にすることが可能
- argparse を利用しても変わらないところ
  - Parse 前に引数の割当を宣言し、Parse 時にそのポインタへ値がセットされるところ
    - python のように、Parse の返却値が Dict 型で返ってくるわけではない
