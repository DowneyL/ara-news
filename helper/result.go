package helper

type AffectNum map[string]int64

func NewAffectNum(v int64) AffectNum {
	af := make(AffectNum)
	af["affect_num"] = v
	return af
}
