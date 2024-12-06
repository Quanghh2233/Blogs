package serializer

import (
	"time"

	"github.com/Quanghh2233/blogs/internal/model"
)

type PostResp struct {
	ID         uint          `json:"id"`
	Title      string        `json:"title"`
	Decription string        `json:"description"`
	Image      string        `json:"image"`
	CreateAt   time.Time     `json:"created_at"`
	UpdateAt   time.Time     `json:"updated_at"`
	Category   CategoryResp  `json:"category"`
	Comment    []CommentResp `json:"comments"`
}

type PostSerializer struct {
	Post model.Post
}

func (serializer *PostSerializer) Response() PostResp {
	categorySerializer := CategorySerializer{Category: serializer.Post.Category}
	CommentSerializer := CommentsSerializer{Comments: serializer.Post.Comments}
	return PostResp{
		ID:         serializer.Post.ID,
		Title:      serializer.Post.Title,
		Decription: serializer.Post.Description,
		Image:      serializer.Post.Image,
		CreateAt:   serializer.Post.CreatedAt,
		UpdateAt:   serializer.Post.UpdatedAt,
		Category:   categorySerializer.Response(),
		Comment:    CommentSerializer.Response(),
	}
}

type PostsSerializer struct {
	Posts []model.Post
}

func (serializer *PostsSerializer) Response() []PostResp {
	response := []PostResp{}
	for _, post := range serializer.Posts {
		postSerializer := PostSerializer{Post: post}
		response = append(response, postSerializer.Response())
	}
	return response
}
