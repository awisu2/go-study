// packageは合わせる
package main

import (
	"testing"
)

func TestDivision(t *testing.T) {
    num, err := Division(9, 3) 
    if num != 3 || err != nil {
        t.Fatalf("Division not correct. a=%d, b=%d, answer=%d", 9, 3, 3)
    }
}

func TestDivisionZero(t *testing.T) {
    num, err := Division(9, 0) 
    if num != 0 || err == nil {
        t.Fatalf("Division(9, 0) want error. %d, %s", num, err)
    }
}