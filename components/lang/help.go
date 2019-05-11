package lang

import "strings"

const (
	zh_CN int = iota + 1
	en_US
)

func GetLangId(lang string) int {
	lang = strings.ToLower(lang)
	switch lang {
	case "zh_cn":
		return zh_CN
	case "en_US":
		return en_US
	default:
		return en_US
	}
}
