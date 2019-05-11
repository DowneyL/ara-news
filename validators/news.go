package validators

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
}
