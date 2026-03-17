package dto

import (
	"time"

	"github.com/google/uuid"
)

type CategoryDTO struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	FormJson  string    `json:"form_json"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryDTO struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	FormJson string `json:"form_json"`
}

type UpdateCategoryDTO struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	FormJson string `json:"form_json"`
}
