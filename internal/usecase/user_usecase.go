package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
	jwtService     jwt.JWTService
}

func NewUserUsecase(ur domain.UserRepository, timeout time.Duration, jwtService jwt.JWTService) domain.UserUsecase {
	return &userUsecase{
		userRepo:       ur,
		contextTimeout: timeout,
		jwtService:     jwtService,
	}
}

func (u *userUsecase) GenerateToken(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) (string, error) {
	return u.jwtService.GenerateToken(id, tenantID)
}

func (u *userUsecase) Login(ctx context.Context, email string, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, domain.ErrUnauthorized
	}

	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	// Check if email already exists
	existedUser, _ := u.userRepo.GetByEmail(ctx, user.Email)
	if existedUser != nil {
		return domain.ErrConflict
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	now := time.Now()
	user.Password = string(hashedPassword)
	user.UUID = uuid.New()
	user.CreatedAt = now
	user.UpdatedAt = &now

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetByID(ctx, id)
}

func (u *userUsecase) GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.GetByUUID(ctx, uuid)
}

func (u *userUsecase) Fetch(ctx context.Context, page int, limit int) ([]domain.User, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	offset := (page - 1) * limit
	return u.userRepo.Fetch(ctx, limit, offset)
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	existingUser, err := u.userRepo.GetByID(ctx, user.UUID)
	if err != nil {
		return err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existingUser.Password = string(hashedPassword)
	}
	now := time.Now()
	existingUser.UpdatedAt = &now

	return u.userRepo.Update(ctx, existingUser)
}

func (u *userUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.Delete(ctx, id)
}
