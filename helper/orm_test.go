package helper

import (
	"fmt"
	"testing"
)

func TestGetOrmOrders(t *testing.T) {
	str := "update DESC, seq"
	strings, e := GetOrmOrders(str)
	fmt.Println(strings)
	fmt.Println(e)
}
