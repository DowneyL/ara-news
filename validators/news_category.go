package validators

type NewsCategory struct {
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
