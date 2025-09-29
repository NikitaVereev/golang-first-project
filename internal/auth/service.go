package auth

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type AuthService struct {
	AuthRepository *AuthRepository
}

func NewAuthService(authRepository *AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}

func (service *AuthService) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
	// Проверяем существующий email
	existing, _ := service.AuthRepository.FindMainRegistryByEmail(ctx, req.Email)
	if existing != nil {
		return nil, errors.New(ErrUserExists)
	}
	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Генерируем уникальный maincabinet
	mainCabinet := service.generateMainCabinet(req.Email)

	// Создаем запись в mainregistry
	registry := &MainRegistry{
		Email:       strings.ToLower(req.Email),
		Phone:       req.Phone,
		Password:    string(hashedPassword),
		Name:        req.Name,
		Position:    req.Position,
		Filials:     req.Filials,
		Brand:       req.Brand,
		MainCabinet: mainCabinet,
	}

	if err := service.AuthRepository.CreateMainRegistry(ctx, registry); err != nil {
		return nil, err
	}

	// Создаем запись в mainusers
	mainUser := &MainUsers{
		Login:       strings.ToLower(req.Email),
		MainCabinet: mainCabinet,
	}

	if err := service.AuthRepository.CreateMainUser(ctx, mainUser); err != nil {
		return nil, err
	}

	// Создаем первого пользователя в кабинете
	authUser := &AuthUser{
		Login:       strings.ToLower(req.Email),
		Email:       strings.ToLower(req.Email),
		Fio:         req.Name,
		Password:    string(hashedPassword),
		IDFather:    "0",
		IDGroup:     "1",
		MainCabinet: mainCabinet,
		Responsible: req.Name,
		Role:        "admin",
	}

	if err := service.AuthRepository.CreateAuthUser(ctx, authUser); err != nil {
		return nil, err
	}

	return &RegisterResponse{
		Status:      "success",
		Message:     "It`s My Life!",
		MainCabinet: mainCabinet,
		UserID:      authUser.ID,
	}, nil
}

func (service *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// Ищем в mainusers чтобы получить maincabinet
	mainUser, err := service.AuthRepository.FindMainUserByLogin(ctx, strings.ToLower(req.Email))
	if err != nil {
		return nil, errors.New(ErrWrongCredentials)
	}

	// Ищем пользователя в кабинете
	authUser, err := service.AuthRepository.FindAuthUserByLogin(ctx, strings.ToLower(req.Email), mainUser.MainCabinet)
	if err != nil {
		return nil, errors.New(ErrWrongCredentials)
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(req.Password)); err != nil {
		return nil, errors.New(ErrWrongCredentials)
	}

	return &LoginResponse{
		Status:      "success",
		MainCabinet: authUser.MainCabinet,
		Role:        authUser.Role,
		UserID:      authUser.ID,
		Email:       authUser.Email,
	}, nil
}

func (s *AuthService) generateMainCabinet(email string) string {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	data := email + timestamp
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)[:16]
}
