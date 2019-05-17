package helper

func RmDuplicate(arr []int64) []int64 {
	resArr := make([]int64, 0)
	tmpMap := make(map[int64]interface{})
	for _, val := range arr {
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = struct{}{}
		}
	}

	return resArr
}
