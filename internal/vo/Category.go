package vo

import "mall-demo/internal/model"

type CategoryResponse struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"categoryName"`
}

func BuildCategoryResponse(category *model.Category) CategoryResponse {
	return CategoryResponse{
		ID:           category.ID,
		CategoryName: category.CategoryName,
	}
}

func BuildCategoryResponseList(categoryList []model.Category) (categoryResponseList []CategoryResponse) {
	for _, category := range categoryList {
		categoryResponse := BuildCategoryResponse(&category)
		categoryResponseList = append(categoryResponseList, categoryResponse)
	}
	return categoryResponseList
}
