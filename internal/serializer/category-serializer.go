package serializer

import "github.com/Quanghh2233/blogs/internal/model"

type CategoryResp struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategorySerializer struct {
	Category model.Category
}

func (serializer *CategorySerializer) Response() CategoryResp {
	return CategoryResp{
		ID:   serializer.Category.ID,
		Name: serializer.Category.Name,
	}
}

type CategoriesSerializer struct {
	Categories []model.Category
}

func (serializer *CategoriesSerializer) Response() []CategoryResp {
	response := []CategoryResp{}
	for _, category := range serializer.Categories {
		categorySerializer := CategorySerializer{Category: category}
		response = append(response, categorySerializer.Response())
	}
	return response
}
