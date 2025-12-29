package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"server/internal/domain"
	"server/internal/usecase"
)

type Handler struct {
	auth *usecase.AuthService
}

func NewHandler(auth *usecase.AuthService) *Handler {
	return &Handler{auth: auth}
}
func NewRouter(handler *Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	return mux
}

type creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var creds creds
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.auth.Register(r.Context(), creds.Username, creds.Password); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			status = http.StatusConflict
		}

		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("register success"))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var creds creds
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.auth.Login(r.Context(), creds.Username, creds.Password); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, domain.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		}

		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login success"))
}
