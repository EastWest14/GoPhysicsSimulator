package converter

import (
	"testing"
)

func TestA(t *testing.T) {
	if num(3) != 3 {
		t.Error("fail!")
	}
}
