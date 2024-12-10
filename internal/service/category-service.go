package service

import (
	"github.com/Quanghh2233/blogs/internal/dto"
	"github.com/Quanghh2233/blogs/internal/model"
	"github.com/Quanghh2233/blogs/internal/repository"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type CategoryService interface {
	All() []model.Category
	Insert(categoryDTO dto.Category) model.Category
	Update(categoryId uint64, categoryDTO dto.Category) (model.Category, error)
	DeleteById(categoryId uint64) *gorm.DB
}

type categoryService struct {
	categoryRepo repository.CategoryRepo
}

func NewCategoryService(categoryRepo repository.CategoryRepo) *categoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (service *categoryService) All() []model.Category {
	return service.categoryRepo.AllCategories()
}

func (service *categoryService) Insert(categoryDTO dto.Category) model.Category {
	categoryModel := model.Category{}
	err := smapping.FillStruct(&categoryModel, smapping.MapFields(&categoryDTO))
	if err != nil {
		panic(err)
	}
	return service.categoryRepo.Insert(categoryModel)
}

func (service *categoryService) Update(categoryId uint64, categoryDTO dto.Category) (model.Category, error) {
	category, err := service.categoryRepo.GetById(categoryId)
	if err != nil {
		return category, err
	}

	fillErr := smapping.FillStruct(&category, smapping.MapFields(&categoryDTO))
	if fillErr != nil {
		panic(fillErr)
	}

	service.categoryRepo.Save(&category)
	return category, nil
}

func (service *categoryService) DeleteById(categoryId uint64) *gorm.DB {
	return service.categoryRepo.DeleteById(categoryId)
}
