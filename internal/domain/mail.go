package domain

import "context"

type MailService interface {
	SendEmail(to string, subject string, body string) error
}

type UserActivation struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Token     string    `json:"token" gorm:"not null"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
	CreatedAt string    `json:"created_at"` // Gunakan time yang sesuai standard proyek
}

type UserActivationRepository interface {
	Create(ctx context.Context, activation *UserActivation) error
	GetByToken(ctx context.Context, token string) (*UserActivation, error)
	Update(ctx context.Context, activation *UserActivation) error
}
