package helper

import "strconv"

func ConvInt64ToInt(raw int64) int {
	formatStr := strconv.FormatInt(raw, 10)
	converted, _ := strconv.Atoi(formatStr)
	return converted
}

func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
