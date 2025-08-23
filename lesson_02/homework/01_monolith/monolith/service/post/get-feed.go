package post

import (
	"context"
	"monolith/domain/post"
	"monolith/domain/user"
)

type GetFeedData struct {
	UserID string
	Offset uint64
	Limit  uint64
}

type GetFeedResult struct {
	Posts []*post.Post
}

type PostGetFeedService struct {
	userRepository user.Repository
	postRepository post.Repository
}

func NewPostGetFeedService(userRepository user.Repository, postRepository post.Repository) *PostGetFeedService {
	return &PostGetFeedService{userRepository: userRepository, postRepository: postRepository}
}

func (s *PostGetFeedService) Handle(ctx context.Context, data *GetFeedData) (*GetFeedResult, error) {
	friends, err := s.userRepository.FindUserFriends(ctx, data.UserID)

	if err != nil {
		return nil, err
	}

	friendsId := s.getFriendsId(friends)

	if len(friendsId) == 0 {
		return &GetFeedResult{Posts: []*post.Post{}}, nil
	}

	posts, err := s.postRepository.Find(ctx, friendsId, data.Offset, data.Limit)

	if err != nil {
		return nil, err
	}

	return &GetFeedResult{Posts: posts}, nil
}

func (s *PostGetFeedService) getFriendsId(friends []*user.User) []string {
	result := []string{}

	for _, friend := range friends {
		result = append(result, friend.ID)
	}

	return result
}
