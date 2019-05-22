package help

type Category struct {
	Content
	Seq      int    `json:"seq" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

type Content struct {
	Lang string `json:"lang" validate:"required"`
	Text string `json:"text" validate:"required,max=200"`
}
