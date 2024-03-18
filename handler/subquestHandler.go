package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
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
	fmt.Fprintf(w, " %d", "subquest is created"+newQuest.Title)
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
