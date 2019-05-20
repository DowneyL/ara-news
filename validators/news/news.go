package news

type News struct {
	Info
	Content
}

type Info struct {
	Author       string `json:"author" validate:"required,max=20"`
	CategoryCode string `json:"category_code" validate:"required,max=50"`
	Platform     string `json:"platform" validate:"required"`
	Seq          int    `json:"seq" validate:"required"`
	IsHidden     bool   `json:"is_hidden" validate:"-"`
	CoverUrl     string `json:"cover_url" validate:"required,max=150"`
	PublishedAt  int64  `json:"published_at" validate:"-"`
}

type Content struct {
	Lang    string `json:"lang" validate:"required"`
	Title   string `json:"title" validate:"required,max=150"`
	Content string `json:"content" validate:"required"`
	Default bool   `json:"default" validate:"-"`
}

type Query struct {
	Ids        []int64 `form:"ids" json:"ids" validate:"-"`
	CategoryId int64   `form:"category_id" json:"category_id" validate:"-"`
	Platform   string  `form:"platform" json:"platform" validate:"-"`
	Author     string  `form:"author" json:"author" validate:"-"`
	OrderBy    string  `form:"author" json:"order_by" validate:"-"`
}
