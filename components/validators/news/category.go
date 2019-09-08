package news

type Category struct {
	Code   string `json:"code" validate:"required,max=50"`
	Seq    int    `json:"seq" validate:"required"`
	Icon   string `json:"icon" validate:"required,max=200,url"`
	NameZH string `json:"name_zh" validate:"max=100"`
	NameEN string `json:"name_en" validate:"required,max=100"`
}

type UpdateNameEn struct {
	NameEN string `json:"name_en" validate:"required,max=100"`
}

type CategoryIds struct {
	Ids []int64 `json:"ids" validate:"gt=0,dive,required"`
}

type QueryCategory struct {
	CIds   []int64 `form:"cids" json:"cids" validate:"-"`
	Code   string  `form:"code" json:"code" validate:"max=50"`
	NameZH string  `form:"name_zh" json:"name_zh" validate:"max=100"`
	NameEN string  `form:"name_en" json:"name_en" validate:"max=100"`
}
