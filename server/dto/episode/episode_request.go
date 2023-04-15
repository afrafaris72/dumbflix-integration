package episodeDto

type CreateEpisodeRequest struct {
	Title     string `json:"title" form:"title" validate:"required"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" validate:"required"`
	VideoLink string `json:"video_link" form:"video_link" validate:"required"`
	MovieID   int    `json:"movie_id" form:"movie_id" validate:"required"`
}

type UpdateEpisodeRequest struct {
	Title     string `json:"title" form:"title"`
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
	VideoLink string `json:"video_link" form:"video_link"`
	MovieID   int    `json:"movie_id" form:"movie_id"`
}
