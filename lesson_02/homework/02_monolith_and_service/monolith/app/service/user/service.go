package user

import "monolith/infrastructure/storage"

type UserService struct {
	GetById      *UserGetByIdService
	Search       *UserSearchService
	AddFriend    *UserAddFriendService
	DeleteFriend *UserDeleteFriendService
}

func NewService(repositories *storage.Repository) *UserService {
	return &UserService{
		GetById:      NewUserGetByIdService(repositories.User),
		Search:       NewUserSearchService(repositories.User),
		AddFriend:    NewUserAddFriendService(repositories.User),
		DeleteFriend: NewUserDeleteFriendService(repositories.User),
	}
}
