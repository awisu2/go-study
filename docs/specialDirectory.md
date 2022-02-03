# 特殊ディレクトリ

特殊ディレクトリの取得([os\-study/specialDirectory](https://github.com/awisu2/os-study/blob/main/common/specialDirectory.md))

```go
package pathes

import (
	"log"
	"os"
)

func GetPath() {
	// C:\Users\{user}\AppData\Roaming <nil>
	log.Println(os.UserConfigDir())
	// C:\Users\{user} <nil>
	log.Println(os.UserHomeDir())
	// C:\Users\{user}\AppData\Local <nil>
	log.Println(os.UserCacheDir())
}
```
