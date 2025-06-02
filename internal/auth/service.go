package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Gitong23/go-fiber-hex-api/internal/user"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// type RegisterResponse struct {
// 	UserID    string `json:"user_id"`
// 	FirstName string `json:"first_name"`
// 	Email     string `json:"email"`
// 	LastName  string `json:"last_name"`
// }

type Service interface {
	Register(req RegisterRequest) (*user.User, error)
	Login(req LoginRequest) (*TokenDetails, error)
}

type service struct {
	// repo      user.UserRepository
	userRepo  user.UserRepository
	jwtSecret string
}

type TokenDetails struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewAuthService(userRepo user.UserRepository, jwtSecret string) Service {
	return &service{
		userRepo,
		jwtSecret,
	}
}

func (s *service) Register(req RegisterRequest) (*user.User, error) {
	ctx := context.Background()

	// Check if user already exists
	// existingUser, err := s.userService.GetUserByUsername(ctx, req.Username)
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

	user := &user.User{
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

func (s *service) Login(req LoginRequest) (*TokenDetails, error) {
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

func (s *service) GenerateToken(user *user.User) (*TokenDetails, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   user.ID.Hex(),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
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
