package main

import (
	"testing"
)

// テスト

func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

func TestAbsSecond(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}

// ベンチマーク
func BenchmarkAbs(b *testing.B) {
    got := Abs(-1)
    if got != 1 {
        b.Errorf("Abs(-1) = %d; want 1", got)
    }
}
