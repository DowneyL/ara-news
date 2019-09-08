package help

type Category struct {
	CateContent
	Seq      int    `json:"seq" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

type CateContent struct {
	Lang string `json:"lang" validate:"required"`
	Text string `json:"text" validate:"required,max=200"`
}
