package core

import (
	"context"
	"errors"
	"time"

	"github.com/Gitong23/go-fiber-hex-api/config"
	userCore "github.com/Gitong23/go-fiber-hex-api/internal/core/user"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IauthService interface {
	Register(req RegisterRequest) (*userCore.User, error)
	Login(req LoginRequest) (*TokenDetails, error)
}

type authService struct {
	userRepo  userCore.IuserRepository
	jwtSecret string
}

type TokenDetails struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewAuthService(userRepo userCore.IuserRepository, jwtSecret string) IauthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(req RegisterRequest) (*userCore.User, error) {
	ctx := context.Background()

	// Check if user already exists
	existingUser, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &userCore.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Remove password from response
	user.Password = ""
	return user, nil
}

func (s *authService) Login(req LoginRequest) (*TokenDetails, error) {
	ctx := context.Background()

	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return s.GenerateToken(user)
}

func (s *authService) GenerateToken(user *userCore.User) (*TokenDetails, error) {
	// Parse expiration duration from config
	expirationDuration, err := time.ParseDuration(config.AppConfig.JWT.ExpiresIn)
	if err != nil {
		// Fallback to 24 hours if parsing fails
		expirationDuration = 24 * time.Hour
	}

	expirationTime := time.Now().Add(expirationDuration)
	claims := &jwt.RegisteredClaims{
		Subject:   user.ID.Hex(),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    config.AppConfig.App.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &TokenDetails{
		Token:     tokenString,
		ExpiresAt: expirationTime,
	}, nil
}
