package helper

import "strconv"

func ConvInt64ToInt(raw int64) int {
	formatStr := strconv.FormatInt(raw, 10)
	converted, _ := strconv.Atoi(formatStr)
	return converted
}
