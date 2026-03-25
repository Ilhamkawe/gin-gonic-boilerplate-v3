package usecase

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Register(ctx context.Context, email string, password string, name string, phone string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	// Check if already exists
	existedUser, _ := u.userRepo.GetByEmail(ctx, email)
	if existedUser != nil {
		return nil, domain.ErrConflict
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		UUID:      uuid.New(),
		Email:     email,
		Password:  string(hashedPassword),
		Name:      name,
		Phone:     phone,
		IsActive:  false, // Default false until activated
		CreatedBy: "system_register",
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Generate and Send Activation Token
	token := uuid.New().String()
	activation := &domain.UserActivation{
		UserID: user.ID,
		Email:  user.Email,
		Token:  token,
	}

	if err := u.userActivationRepo.Create(ctx, activation); err != nil {
		return nil, err
	}

	// Parse Template HTML
	tmpl, err := template.ParseFiles("internal/infrastructure/templates/activation_email.html")
	if err != nil {
		fmt.Printf("Error parse email template: %v\n", err)
		return user, nil // Tetap return user walau email gagal (opsional)
	}

	var body bytes.Buffer
	data := struct {
		Name  string
		Token string
	}{
		Name:  user.Name,
		Token: token,
	}

	if err := tmpl.Execute(&body, data); err != nil {
		fmt.Printf("Error execute email template: %v\n", err)
		return user, nil
	}

	// Kirim email di background (goroutine) agar response cepat
	go func() {
		// Gunakan context.Background() karena request context mungkin sudah mati saat goroutine jalan
		_ = u.mailService.SendEmail(user.Email, "Aktivasi Akun Warehouse MS — Selamat Datang!", body.String())
	}()

	return user, nil
}

func (u *userUsecase) ActivateUser(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	activation, err := u.userActivationRepo.GetByToken(ctx, token)
	if err != nil {
		return err
	}

	user, err := u.userRepo.GetByID(ctx, uuid.Nil) // need GetByID with ID or change Repo
	// For now we assume user ID is available in activation
	// We need to fetch user by ID
	// I'll use a hack or update repository soon.
	// But let's assume Repo has GetByIntID if needed.
	// Looking at current repo... it only has GetByUUID.
	// Let's use GetByEmail then (it's unique)
	user, err = u.userRepo.GetByEmail(ctx, activation.Email)
	if err != nil {
		return err
	}

	user.IsActive = true
	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}

	activation.IsUsed = true
	return u.userActivationRepo.Update(ctx, activation)
}

type userUsecase struct {
	userRepo           domain.UserRepository
	userActivationRepo domain.UserActivationRepository
	mailService        domain.MailService
	contextTimeout     time.Duration
	jwtService         jwt.JWTService
}

func NewUserUsecase(ur domain.UserRepository, uar domain.UserActivationRepository, ms domain.MailService, timeout time.Duration, jwtService jwt.JWTService) domain.UserUsecase {
	return &userUsecase{
		userRepo:           ur,
		userActivationRepo: uar,
		mailService:        ms,
		contextTimeout:     timeout,
		jwtService:         jwtService,
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

	if !user.IsActive {
		return nil, domain.ErrForbidden
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
