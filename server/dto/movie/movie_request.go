package moviesDto

type CreateMovieRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail" validate:"required"`
	Year        string `json:"year" form:"year" validate:"required"`
	CategoryID  int    `json:"category_id" form:"category_id" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

type UpdateMovieRequest struct {
	Title       string `json:"title" form:"title"`
	Thumbnail   string `json:"thumbnail" form:"thumbnail"`
	Year        string `json:"year" form:"year"`
	CategoryID  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
}
