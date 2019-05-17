package helper

import (
	"fmt"
	"testing"
)

func TestRmDuplicate(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 22, 5, 44, 3}
	arr = RmDuplicate(arr)
	arr2 := make([]int64, 0)
	arr2 = RmDuplicate(arr2)
	fmt.Println(arr)
	fmt.Println(arr2)
}
