package schemas

import "github.com/google/uuid"

type CreateBlogRequest struct {
	Text string `json:"text" binding:"required"`
}

type UpdateBlogRequest struct {
	Text string `json:"text"`
}

type BlogResponse struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}