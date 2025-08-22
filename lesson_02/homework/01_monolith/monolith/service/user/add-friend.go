package user

import (
	"context"
	"errors"
	"monolith/domain/user"
)

type AddFriendRequest struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
}

type UserAddFriendService struct {
	userRepository user.Repository
}

func NewUserAddFriendService(userRepository user.Repository) *UserAddFriendService {
	return &UserAddFriendService{userRepository: userRepository}
}

func (s *UserAddFriendService) Handle(ctx context.Context, data AddFriendRequest) error {
	usersAreFriends, err := s.userRepository.CheckIfUsersAreFriends(ctx, data.UserID, data.FriendID)

	if err != nil {
		return err
	}

	if usersAreFriends {
		return errors.New("users are already friends")
	}

	if err = s.userRepository.AddFriend(ctx, data.UserID, data.FriendID); err != nil {
		return err
	}

	return nil
}
