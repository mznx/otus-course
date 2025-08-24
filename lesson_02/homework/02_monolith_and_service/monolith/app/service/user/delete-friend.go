package user

import (
	"context"
	"errors"
	"monolith/domain/user"
)

type DeleteFriendData struct {
	UserID   string
	FriendID string
}

type UserDeleteFriendService struct {
	userRepository user.Repository
}

func NewUserDeleteFriendService(userRepository user.Repository) *UserDeleteFriendService {
	return &UserDeleteFriendService{userRepository: userRepository}
}

func (s *UserDeleteFriendService) Handle(ctx context.Context, data *DeleteFriendData) error {
	usersAreFriends, err := s.userRepository.CheckIfUsersAreFriends(ctx, data.UserID, data.FriendID)

	if err != nil {
		return err
	}

	if !usersAreFriends {
		return errors.New("users are not friends")
	}

	if err = s.userRepository.DeleteFriend(ctx, data.UserID, data.FriendID); err != nil {
		return err
	}

	return nil
}
