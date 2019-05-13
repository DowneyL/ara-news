package lang

type Lid int

const (
	zhCN Lid = iota + 1
	enUS
	ja
)

const (
	zhCNCode string = "zh-CN"
	enUSCode        = "en-US"
	jaCode          = "ja"
)

func GetLangId(lang string) Lid {
	switch lang {
	case zhCNCode:
		return zhCN
	case enUSCode:
		return enUS
	case jaCode:
		return ja
	default:
		return enUS
	}
}

func (l *Lid) String() string {
	switch *l {
	case zhCN:
		return zhCNCode
	case enUS:
		return enUSCode
	case ja:
		return jaCode
	default:
		return enUSCode
	}
}
