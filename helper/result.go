package helper

type NumResult map[string]int64

func newNumResult(k string, v int64) NumResult {
	nr := make(NumResult)
	nr[k] = v
	return nr
}

func NewAffectNum(v int64) NumResult {
	return newNumResult("affect_num", v)
}

func NewInsertId(v int64) NumResult {
	return newNumResult("insert_id", v)
}
