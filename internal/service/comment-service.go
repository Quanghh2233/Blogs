package service

import (
	"github.com/Quanghh2233/blogs/internal/dto"
	"github.com/Quanghh2233/blogs/internal/model"
	"github.com/Quanghh2233/blogs/internal/repository"
	"github.com/mashingan/smapping"
)

type CommentService interface {
	All(limit string, offset string, postId uint) []model.Comment
	Insert(commentDto dto.Comment, postId, userId uint) model.Comment
}

type commentService struct {
	commentRepo repository.CommentRepo
}

func NewCommentService(commentRepo repository.CommentRepo) *commentService {
	return &commentService{commentRepo: commentRepo}
}

func (service *commentService) All(limit string, offset string, postId uint) []model.Comment {
	return service.commentRepo.AllCommentByPostId(limit, offset, postId)
}

func (service *commentService) Insert(commentDto dto.Comment, postId, userId uint) model.Comment {
	commentModel := model.Comment{}
	err := smapping.FillStruct(&commentModel, smapping.MapFields(&commentDto))
	if err != nil {
		panic(err)
	}
	commentModel.PostID = postId
	commentModel.UserID = userId
	return service.commentRepo.Insert(commentModel)
}
