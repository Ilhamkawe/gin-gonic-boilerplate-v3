package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userUUID uuid.UUID, tenantID uuid.UUID) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey []byte
	issuer    string
}

func NewJWTService(secretKey string) *jwtService {
	return &jwtService{
		secretKey: []byte(secretKey),
		issuer:    "warehouse.api",
	}
}

type AuthCustomClaims struct {
	UserUUID   uuid.UUID `json:"user_uuid"`
	TenantUUID uuid.UUID `json:"tenant_uuid"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(userUUID uuid.UUID, tenantUUID uuid.UUID) (string, error) {
	claims := &AuthCustomClaims{
		UserUUID:   userUUID,
		TenantUUID: tenantUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})
}

func (j *jwtService) GetClaims(tokenString string) (*AuthCustomClaims, error) {
	token, err := j.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	return token.Claims.(*AuthCustomClaims), nil
}
