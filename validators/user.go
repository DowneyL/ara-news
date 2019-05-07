package validators

type Device struct {
	Token string `form:"token" validate:"required"`
}

type UserRegister struct {
	Name string `form:"name" validate:"required,min=5,max=30"`
}
