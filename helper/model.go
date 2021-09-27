package helper

import (
	"fandy123/belajar-golang-restful-api/model/domain"
	"fandy123/belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Catagory) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Catagory) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
