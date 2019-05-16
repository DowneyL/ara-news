package lang

type Lid int

const (
	ZhCN Lid = iota + 1
	EnUS
	Ja
)

const (
	ZhCNCode string = "zh-CN"
	EnUSCode        = "en-US"
	JaCode          = "ja"
)

func GetLangId(lang string) Lid {
	switch lang {
	case ZhCNCode:
		return ZhCN
	case EnUSCode:
		return EnUS
	case JaCode:
		return Ja
	default:
		return EnUS
	}
}

func (l *Lid) String() string {
	switch *l {
	case ZhCN:
		return ZhCNCode
	case EnUS:
		return EnUSCode
	case Ja:
		return JaCode
	default:
		return EnUSCode
	}
}
