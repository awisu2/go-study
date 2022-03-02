# callback study

## samples

```go
package main

func callback(f func()) {
	f()
}

// closure
func counter(f func(i int)) func(j int) {
	count := 0
	return func(j int) {
		count += j
		f(count)
	}
}
```


```go
package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	callback := func(i int) {
		if i != 1 && i != 3 {
			t.Errorf("got %v. not target value", i)
		}
	}
	countUp := counter(callback)

	countUp(1)
	countUp(2)
}
```
