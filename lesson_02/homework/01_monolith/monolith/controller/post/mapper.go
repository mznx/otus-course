package post

import (
	"encoding/json"
	"monolith/helper"
	post_service "monolith/service/post"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GenerateCreatePostData(r *http.Request) (*post_service.PostCreateData, error) {
	var body CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &post_service.PostCreateData{
		UserID: helper.GetAuthorizedUserId(r),
		Text:   body.Text,
	}, nil
}

func GenerateUpdatePostData(r *http.Request) (*post_service.UpdateData, error) {
	var body UpdatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return &post_service.UpdateData{
		PostID: body.PostID,
		Text:   body.Text,
	}, nil
}

func GenerateDeletePostData(r *http.Request) (*post_service.DeleteData, error) {
	return &post_service.DeleteData{
		PostID: chi.URLParam(r, "post_id"),
	}, nil
}

func GenerateGetByIdtData(r *http.Request) (*post_service.GetByIdData, error) {
	return &post_service.GetByIdData{
		PostID: chi.URLParam(r, "post_id"),
	}, nil
}

func GenerateGetByIdtResponse(result *post_service.GetByIdResult) *GetByIdResponse {
	return &GetByIdResponse{
		ID:       result.Post.ID,
		Text:     result.Post.Text,
		AuthorID: result.Post.AuthorID,
	}
}

func GenerateGetFeedData(r *http.Request) (*post_service.GetFeedData, error) {
	offset := uint64(0)
	limit := uint64(100)

	queryOffset := r.URL.Query().Get("offset")
	queryLimit := r.URL.Query().Get("limit")

	if queryOffset != "" {
		parsedOffset, err := strconv.ParseUint(queryOffset, 10, 32)

		if err != nil {
			return nil, err
		}

		offset = parsedOffset
	}

	if queryLimit != "" {
		parsedLimit, err := strconv.ParseUint(queryLimit, 10, 32)

		if err != nil {
			return nil, err
		}

		limit = parsedLimit
	}

	return &post_service.GetFeedData{
		UserID: helper.GetAuthorizedUserId(r),
		Offset: offset,
		Limit:  limit,
	}, nil
}

func GenerateGetFeedResponse(result *post_service.GetFeedResult) []*GetFeedResponse {
	response := []*GetFeedResponse{}

	for _, post := range result.Posts {
		response = append(response, &GetFeedResponse{
			ID:       post.ID,
			Text:     post.Text,
			AuthorID: post.AuthorID,
		})
	}

	return response
}
