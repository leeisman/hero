package utils

import (
	"fmt"
	"testing"
)

func TestGenMd5Key(t *testing.T) {
	got := GenMd5Key("3299005360117810")
	fmt.Println("!!!!!!!!!!!!!!", got)
}
