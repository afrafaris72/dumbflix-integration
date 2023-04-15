package categoryDto

type CreateCategoryRequest struct {
	Name string `Json:"name" form:"name" validate:"required"`
}

type UpdateCategoryRequest struct {
	Name string `Json:"name" form:"name"`
}
