package auth

import (
	"context"
	"first_project/pkg/db"
)

type Repository interface {
	// Регистрация
	CreateMainRegistry(ctx context.Context, registry *MainRegistry) error
	CreateMainUser(ctx context.Context, mainUser *MainUsers) error
	CreateAuthUser(ctx context.Context, authUser *AuthUser) error

	// Авторизация
	FindMainUserByLogin(ctx context.Context, login string) (*MainUsers, error)
	FindAuthUserByLogin(ctx context.Context, login, mainCabinet string) (*AuthUser, error)
	FindMainRegistryByEmail(ctx context.Context, email string) (*MainRegistry, error)
}

type AuthRepository struct {
	db *db.Db
}

func NewAuthRepository(db *db.Db) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateMainRegistry(ctx context.Context, registry *MainRegistry) error {
	result := r.db.WithContext(ctx).Create(registry)
	return result.Error
}

func (r *AuthRepository) CreateMainUser(ctx context.Context, mainUser *MainUsers) error {
	result := r.db.WithContext(ctx).Create(mainUser)
	return result.Error
}

func (r *AuthRepository) CreateAuthUser(ctx context.Context, authUser *AuthUser) error {
	result := r.db.WithContext(ctx).Create(authUser)
	return result.Error
}

func (r *AuthRepository) FindMainUserByLogin(ctx context.Context, login string) (*MainUsers, error) {
	var mainUser MainUsers
	result := r.db.WithContext(ctx).Where("login = ?", login).First(&mainUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &mainUser, nil
}

func (r *AuthRepository) FindAuthUserByLogin(ctx context.Context, login, mainCabinet string) (*AuthUser, error) {
	var authUser AuthUser
	result := r.db.WithContext(ctx).
		Where("login = ? AND maincabinet = ?", login, mainCabinet).
		First(&authUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &authUser, nil
}

func (r *AuthRepository) FindMainRegistryByEmail(ctx context.Context, email string) (*MainRegistry, error) {
	var registry MainRegistry
	result := r.db.WithContext(ctx).Where("email = ?", email).First(&registry)
	if result.Error != nil {
		return nil, result.Error
	}
	return &registry, nil
}
