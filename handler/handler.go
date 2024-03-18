package handler

import (
	"net/http"
	"vk/internal/usecase"
)

type Handler struct {
	service *usecase.UseCase
}

func NewHandler(service *usecase.UseCase) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/quest/user", http.HandlerFunc(h.userHandler))
	mux.Handle("/quest/user/", http.HandlerFunc(h.userHandler))
	mux.Handle("/quest/quest", http.HandlerFunc(h.questHandler))
	mux.Handle("/quest/subquest", http.HandlerFunc(h.subquestHandler))
	mux.Handle("/quest/subquest/", http.HandlerFunc(h.subquestHandler))
	return mux
}
