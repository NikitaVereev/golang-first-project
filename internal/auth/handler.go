package auth

import (
	"errors"
	"first_project/configs"
	"first_project/pkg/jwt"
	"first_project/pkg/req"
	"first_project/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandlerBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		response, err := handler.AuthService.Register(r.Context(), body)
		if err != nil {
			if err == errors.New(ErrUserExists) {
				res.Json(w, map[string]string{
					"status":  "error",
					"message": "User already exists",
				}, http.StatusConflict)
			} else {
				res.Json(w, map[string]string{
					"status":  "error",
					"message": err.Error(),
				}, http.StatusInternalServerError)
			}
			return
		}
		data := RegisterResponse{
			Status:      "success",
			Message:     "Registration successful",
			MainCabinet: response.MainCabinet,
			UserID:      response.UserID,
		}
		res.Json(w, data, http.StatusCreated)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandlerBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		response, err := handler.AuthService.Login(r.Context(), body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email:       response.Email,
			UserID:      fmt.Sprintf("%d", response.UserID),
			MainCabinet: response.MainCabinet,
			Role:        response.Role,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			Status:      "success",
			Token:       token,
			MainCabinet: response.MainCabinet,
			Role:        response.Role,
			UserID:      response.UserID,
			Email:       response.Email,
		}

		res.Json(w, data, http.StatusOK)
	}
}
