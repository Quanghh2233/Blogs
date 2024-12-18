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

type CategoryController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	DeleteById(context *gin.Context)
}

type categoryController struct {
	categoryService service.CategoryService
	logger          *logger.Logger
}

func NewCategoryController(categoryService service.CategoryService, logger *logger.Logger) *categoryController {
	return &categoryController{categoryService: categoryService, logger: logger}
}

// GetCategories             godoc
// @Summary      Get categories list
// @Description  Responds with the list of all categories as JSON.
// @Tags         categories
// @Produce      json
// @Success      200  {object}  serializer.CategoryResp
// @Router       /categories [get]
func (controller *categoryController) All(context *gin.Context) {
	categories := controller.categoryService.All()
	serializer := serializer.CategoriesSerializer{Categories: categories}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// InsertCategory             godoc
// @Summary      Insert category
// @Description  Responds with category as JSON.
// @Tags         categories
// @Produce      json
// @Param data body dto.Category true "Category dto"
// @Success      201  {object}  serializer.CategoryResp
// @Router       /categories [post]
func (controller *categoryController) Insert(context *gin.Context) {
	categoryDto := dto.Category{}
	err := context.ShouldBindJSON(&categoryDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	category := controller.categoryService.Insert(categoryDto)
	serializer := serializer.CategorySerializer{Category: category}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// UpdateCategory             godoc
// @Summary      Update category
// @Description  Responds with category as JSON.
// @Tags         categories
// @Produce      json
// @Param        id  path      uint  true  "update category by id"
// @Param data body dto.Category true "Category dto"
// @Success      200  {object}  serializer.CategoryResp
// @Router       /categories/{id} [put]
func (controller *categoryController) Update(context *gin.Context) {
	categoryDto := dto.Category{}
	err := context.ShouldBindJSON(&categoryDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	categortId, _ := strconv.ParseUint(context.Param("categoryId"), 10, 64)
	category, err := controller.categoryService.Update(categortId, categoryDto)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.ErrorsResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	serializer := serializer.CategorySerializer{Category: category}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// DeleteCategory             godoc
// @Summary      Delete category
// @Description  Responds with category as JSON.
// @Tags         categories
// @Produce      json
// @Param        id  path      uint  true  "delete category by id"
// @Success      204
// @Router       /categories/{id} [delete]
func (controller *categoryController) DeleteById(context *gin.Context) {
	categortId, _ := strconv.ParseUint(context.Param("categoryId"), 10, 64)
	result := controller.categoryService.DeleteById(categortId)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, utils.ErrorsResponse(result.Error.Error()))
		controller.logger.Error().Err(result.Error).Msg("")
		return
	} else if result.RowsAffected < 1 {
		context.JSON(http.StatusNotFound, utils.ErrorsResponse("category does not exists"))
		controller.logger.Error().Msg("category does not exist")
		return
	}
	context.JSON(http.StatusNoContent, utils.GetResponse(gin.H{}))
}
