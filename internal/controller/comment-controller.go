package controller

import (
	"net/http"
	"strconv"

	"github.com/Quanghh2233/blogs/internal/dto"
	"github.com/Quanghh2233/blogs/internal/serializer"
	"github.com/Quanghh2233/blogs/internal/service"
	"github.com/Quanghh2233/blogs/internal/utils"
	"github.com/Quanghh2233/blogs/pkg/logger"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
}

type commentController struct {
	commentService service.CommentService
	logger         *logger.Logger
}

func NewCommentController(commentService service.CommentService, logger *logger.Logger) *commentController {
	return &commentController{
		commentService: commentService,
		logger:         logger,
	}
}

// GetComments             godoc
// @Summary      Get comments list by postId
// @Description  Responds with the list of all comments by postId as JSON.
// @Tags         comments
// @Produce      json
// @Success      200  {object}  serializer.CommentResp
// @Router       /posts/{postId}/comments [get]
func (controller *commentController) All(context *gin.Context) {
	limit := context.Query("limit")
	offset := context.Query("offset")
	postId, err := strconv.ParseUint(context.Param("postId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse("postId param not found"))
		return
	}
	comments := controller.commentService.All(limit, offset, uint(postId))
	serializer := serializer.CommentsSerializer{Comments: comments}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// InsertComment             godoc
// @Summary      Insert comment
// @Description  Responds with comment as JSON.
// @Tags         comments
// @Produce      json
// @Param        postId  path      uint  true  "Insert comment by postId"
// @Param data body dto.Comment true "Comment dto"
// @Success      201  {object}  serializer.CommentResp
// @Router       /posts/{postId}/comments [post]
func (controller *commentController) Insert(context *gin.Context) {
	commentDto := dto.Comment{}
	err := context.ShouldBindJSON(&commentDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	postId, err := strconv.ParseUint(context.Param("postId"), 10, 32)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse("postId param not founf"))
		controller.logger.Error().Msg("postId param not founf")
		return
	}

	tokenString := utils.GetTokenString(context)
	userId, err := utils.GetUserIDFromToken(tokenString)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse("Failed to get userId from token"))
		controller.logger.Error().Msg("Failed to get userId from token")
		return
	}

	comment := controller.commentService.Insert(commentDto, uint(postId), userId)
	serializer := serializer.CommentSerializer{Comment: comment}
	context.JSON(http.StatusCreated, utils.GetResponse(serializer.Response()))
}
