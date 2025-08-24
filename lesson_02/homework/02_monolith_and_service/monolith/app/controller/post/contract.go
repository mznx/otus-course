package post

type CreatePostRequest struct {
	Text string `json:"text"`
}

type UpdatePostRequest struct {
	PostID string `json:"id"`
	Text   string `json:"text"`
}

type GetByIdResponse struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	AuthorID string `json:"author_user_id"`
}

type GetFeedResponse struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	AuthorID string `json:"author_user_id"`
}
