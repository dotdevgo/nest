package dto

type FormTemplate struct {
	Name string `form:"name" validate:"required"`
}
