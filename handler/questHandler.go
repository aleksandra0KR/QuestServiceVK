package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vk/internal/model/quest"
)

func (h *Handler) questHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request on %s", r.Method, r.RequestURI)

	switch r.Method {
	case http.MethodPost:
		h.createQuest(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createQuest(w http.ResponseWriter, r *http.Request) {
	var quest quest.Quest
	err := json.NewDecoder(r.Body).Decode(&quest)
	fmt.Println(quest.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	newQuest, err := h.service.QuestUseCase.Create(&quest)
	fmt.Println(quest.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("createQuest is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%d", "quest is created"+newQuest.Title)
}
