package help

type Content struct {
	Cid      int64  `json:"cid" validate:"required"`
	Lang     string `json:"lang" validate:"required"`
	Question string `json:"question" validate:"required,max=200"`
	Answer   string `json:"answer" validate:"required"`
	Seq      int    `json:"seq" validate:"required"`
}

type Query struct {
	Platform string `form:"platform" json:"platform"`
	OrderBy  string `form:"order_by" json:"order_by"`
}
