package validators

type User struct {
	Name string `form:"name" validate:"required"`
}
