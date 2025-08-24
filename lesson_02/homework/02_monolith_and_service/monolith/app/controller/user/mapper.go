package user

import (
	"monolith/helper"
	user_service "monolith/service/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GenerateGetByIdData(r *http.Request) (*user_service.GetByIdData, error) {
	return &user_service.GetByIdData{
		UserID: chi.URLParam(r, "user_id"),
	}, nil
}

func GenerateGetByIdResponse(result *user_service.GetByIdResult) *GetByIdResponse {
	return &GetByIdResponse{
		UserID:     result.User.ID,
		FirstName:  result.User.FirstName,
		SecondName: result.User.SecondName,
	}
}

func GenerateSearchData(r *http.Request) (*user_service.SearchData, error) {
	return &user_service.SearchData{
		FirstName:  r.URL.Query().Get("first_name"),
		SecondName: r.URL.Query().Get("last_name"),
	}, nil
}

func GenerateSearchResponse(result *user_service.SearchResult) []*SearchResponse {
	response := []*SearchResponse{}

	for _, u := range result.Users {
		response = append(response, &SearchResponse{
			UserID:     u.ID,
			FirstName:  u.FirstName,
			SecondName: u.SecondName,
		})
	}

	return response
}

func GenerateAddFriendData(r *http.Request) (*user_service.AddFriendData, error) {
	return &user_service.AddFriendData{
		UserID:   helper.GetAuthorizedUserId(r),
		FriendID: chi.URLParam(r, "user_id"),
	}, nil
}

func GenerateDeleteFriendData(r *http.Request) (*user_service.DeleteFriendData, error) {
	return &user_service.DeleteFriendData{
		UserID:   helper.GetAuthorizedUserId(r),
		FriendID: chi.URLParam(r, "user_id"),
	}, nil
}
