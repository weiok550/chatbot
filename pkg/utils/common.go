package utils

import "time"

func Now() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func StringIsEmpty(str string) bool {
	return len(str) == 0
}