package util

import (
	"fmt"
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	s := []int{2, 4, 5, 6, 7, 8}
	fmt.Println(Contains(s, 5))
}
