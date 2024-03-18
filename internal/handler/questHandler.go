package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
	"vk/internal/model/quest"
)

func (h *Handler) questHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request on %s", r.Method, r.RequestURI)

	switch r.Method {
	case http.MethodPost:
		h.createQuest(w, r)
	case http.MethodDelete:
		h.deleteQuest(w, r)
	case http.MethodPut:
		if (regexp.MustCompile(`/quest/quest/updatetitle/*`)).MatchString(r.URL.RequestURI()) {
			h.updateTitleQuest(w, r)
		} else if (regexp.MustCompile(`/quest/quest/updatedescription/*`)).MatchString(r.URL.RequestURI()) {
			h.updateDescriptionQuest(w, r)
		} else if (regexp.MustCompile(`/quest/quest/updateduedate/*`)).MatchString(r.URL.RequestURI()) {
			h.updateDueDateQuest(w, r)
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createQuest(w http.ResponseWriter, r *http.Request) {
	var quest quest.Quest
	err := json.NewDecoder(r.Body).Decode(&quest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	newQuest, err := h.service.QuestUseCase.Create(&quest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("createQuest is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "quest is created "+newQuest.Title)
}

func (h *Handler) deleteQuest(w http.ResponseWriter, r *http.Request) {
	id, err := getUuidFromRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.QuestUseCase.DeleteQuestByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("deleteQuest is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "quest is deleted ")
}

func (h *Handler) updateTitleQuest(w http.ResponseWriter, r *http.Request) {
	id, title, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.QuestUseCase.ChangeTitle(id, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateTitleQuest is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "quests titles was changed")
}

func (h *Handler) updateDescriptionQuest(w http.ResponseWriter, r *http.Request) {
	id, description, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.QuestUseCase.ChangeDescription(id, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateDescriptionQuest is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "quests description was changed")
}

func (h *Handler) updateDueDateQuest(w http.ResponseWriter, r *http.Request) {
	id, dueDateStr, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	dueDate, err := time.Parse("2006-01-02T15:04:05Z07:00", dueDateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	err = h.service.QuestUseCase.ChangeDueDate(id, dueDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateDueDateQuest is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "quests due date was changed")
}
