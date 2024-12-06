package serializer

import "github.com/Quanghh2233/blogs/internal/model"

type CommentResp struct {
	ID   uint   `json:"id"`
	Body string `json:"body"`
}

type CommentSerializer struct {
	Comment model.Comment
}

func (serializer *CommentSerializer) Response() CommentResp {
	return CommentResp{
		ID:   serializer.Comment.ID,
		Body: serializer.Comment.Body,
	}
}

type CommentsSerializer struct {
	Comments []model.Comment
}

func (serializer *CommentsSerializer) Response() []CommentResp {
	response := []CommentResp{}
	for _, comment := range serializer.Comments {
		commentSerializer := CommentSerializer{Comment: comment}
		response = append(response, commentSerializer.Response())
	}
	return response
}
