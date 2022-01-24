package dto

// FormTemplate godoc
type FormTemplate struct {
	UUID string `form:"id" json:"id"`
	Name string `form:"name" json:"name" validate:"required"`
}
