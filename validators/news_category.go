package validators

type NewsCategory struct {
	Code   string `json:"code" validate:"required,max=50"`
	Seq    int    `json:"seq"`
	Icon   string `json:"icon" validate:"required,max=200,url"`
	NameZH string `json:"name_zh" validate:"max=100"`
	NameEN string `json:"name_en" validate:"required,max=100"`
}
