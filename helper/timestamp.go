package helper

import (
	"reflect"
	"time"
)

type Timestamp int64

func NewTimestamp(unix ...int64) Timestamp {
	var (
		uTime     int64
		timestamp Timestamp
	)

	if len(unix) > 0 && unix[0] != 0 {
		uTime = unix[0]
	} else {
		uTime = time.Now().Unix()
	}
	reflect.ValueOf(&timestamp).Elem().SetInt(uTime)

	return timestamp
}

func (t *Timestamp) String() string {
	intTime := reflect.ValueOf(*t).Int()

	return Date("Y-m-d H:i:s", intTime)
}
