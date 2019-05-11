package helper

import "strings"

const (
	iOS     int = 1
	android int = 3
)

func GetAttrSetId(platform string) int {
	platform = strings.ToLower(platform)
	switch platform {
	case "ios":
		return iOS
	case "android":
		return android
	default:
		return iOS
	}
}
