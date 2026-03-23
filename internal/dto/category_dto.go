package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/datatypes"
)

type CategoryDTO struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	FormJson  string    `json:"form_json"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	FormJson  string    `json:"form_json"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategory struct {
	Name     string         `form:"name"`
	Icon     string         `form:"icon"`
	Tagline  string         `form:"tagline"`
	FormJson datatypes.JSON `form:"form_json"`
	TenantID int            `header:"X-Tenant-ID"`
}

type UpdateCategory struct {
	UUID     uuid.UUID      `form:"uuid"`
	Name     string         `form:"name"`
	Icon     string         `form:"icon"`
	Tagline  string         `form:"tagline"`
	FormJson datatypes.JSON `form:"form_json"`
}

type InsightCategoryDTO struct {
	ActiveCategories   int64 `json:"active_categories"`
	TotalCategories    int64 `json:"total_categories"`
	InactiveCategories int64 `json:"inactive_categories"`
}

func FromCategory(category domain.Category) CategoryResponse {
	categoryResponse := CategoryResponse{
		UUID:      category.UUID,
		Name:      category.Name,
		Icon:      category.Icon,
		FormJson:  string(category.FormJson),
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
	return categoryResponse
}

func FromCategories(categories []domain.Category) []CategoryResponse {
	categoryResponses := make([]CategoryResponse, 0)
	for _, category := range categories {
		categoryResponses = append(categoryResponses, FromCategory(category))
	}
	return categoryResponses
}

func FromInsightCategory(insight domain.InsightCategory) InsightCategoryDTO {
	return InsightCategoryDTO{
		ActiveCategories:   insight.ActiveCategories,
		TotalCategories:    insight.TotalCategories,
		InactiveCategories: insight.InactiveCategories,
	}
}
