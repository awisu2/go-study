package sub

import (
	"testing"
)

// テスト

func TestHello(t *testing.T) {
    got := Hello("go")
    if got != "hello gao" {
        t.Errorf("Hello() = %s; want hello go", got)
    }
}

// ベンチマーク
func BenchmarkHello(b *testing.B) {
    got := Hello("go")
    if got != "hello go" {
        b.Errorf("ello() = %s; want hello go", got)
    }
}
