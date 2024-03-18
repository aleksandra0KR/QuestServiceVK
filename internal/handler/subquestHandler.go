package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
	"vk/internal/model/subquest"
)

func (h *Handler) subquestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request on %s", r.Method, r.RequestURI)
	switch r.Method {
	case http.MethodPost:
		log.Printf("%s request on %s", r.Method, r.RequestURI)
		if (regexp.MustCompile(`/quest/subquest*$`)).MatchString(r.URL.RequestURI()) {
			h.createSubquest(w, r)
		} else if (regexp.MustCompile(`/quest/subquest/attach/*`)).MatchString(r.URL.RequestURI()) {
			h.attachSubquest(w, r)
		}
	case http.MethodDelete:
		h.deleteSubquest(w, r)
	case http.MethodPut:
		if (regexp.MustCompile(`/quest/subquest/updatetitle/*`)).MatchString(r.URL.RequestURI()) {
			h.updateTitle(w, r)
		} else if (regexp.MustCompile(`/quest/subquest/updatedescription/*`)).MatchString(r.URL.RequestURI()) {
			h.updateDescription(w, r)
		} else if (regexp.MustCompile(`/quest/subquest/updateduedate/*`)).MatchString(r.URL.RequestURI()) {
			h.updateDueDate(w, r)
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createSubquest(w http.ResponseWriter, r *http.Request) {
	var subquest subquest.Subquest
	err := json.NewDecoder(r.Body).Decode(&subquest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	newQuest, err := h.service.SubquestUseCase.Create(&subquest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("createSubquest is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "subquest is created"+newQuest.Title)
}

func (h *Handler) attachSubquest(w http.ResponseWriter, r *http.Request) {
	idQuest, idSubquest, err := getTwoUuidFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.SubquestUseCase.AttachToQuest(idQuest, idSubquest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("attachSubquest is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "subquest is attached")
}

func (h *Handler) deleteSubquest(w http.ResponseWriter, r *http.Request) {
	id, err := getUuidFromRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.SubquestUseCase.DeleteSubquestByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("deleteSubquest is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "subquest is deleted ")
}

func (h *Handler) updateTitle(w http.ResponseWriter, r *http.Request) {
	id, title, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.SubquestUseCase.ChangeTitle(id, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateTitle is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "subquests titles was changed")
}

func (h *Handler) updateDescription(w http.ResponseWriter, r *http.Request) {
	id, description, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.SubquestUseCase.ChangeDescription(id, description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateDescription is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "subquests description was changed")
}

func (h *Handler) updateDueDate(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.SubquestUseCase.ChangeDueDate(id, dueDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateDueDate is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "subquests due date was changed")
}
