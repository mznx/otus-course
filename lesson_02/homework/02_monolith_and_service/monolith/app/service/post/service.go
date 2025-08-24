package post

import "monolith/infrastructure/storage"

type PostService struct {
	CreatePost *PostCreateService
	UpdatePost *PostUpdateService
	DeletePost *PostDeleteService
	GetById    *PostGetByIdService
	GetFeed    *PostGetFeedService
}

func NewService(repositories *storage.Repository) *PostService {
	return &PostService{
		CreatePost: NewPostCreateService(repositories.Post),
		UpdatePost: NewPostUpdateService(repositories.Post),
		DeletePost: NewPostDeleteService(repositories.Post),
		GetById:    NewPostGetByIdService(repositories.Post),
		GetFeed:    NewPostGetFeedService(repositories.User, repositories.Post),
	}
}
