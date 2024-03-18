package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"vk/internal/model/status"
	"vk/internal/model/user"
)

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request on %s", r.Method, r.RequestURI)

	switch r.Method {
	case http.MethodPost:
		if (regexp.MustCompile(`/quest/user/takequest/*`)).MatchString(r.URL.RequestURI()) {
			h.attachUser(w, r)
		} else if (regexp.MustCompile(`/quest/user/takesubquest/*`)).MatchString(r.URL.RequestURI()) {
			h.attachUserSubquest(w, r)
		} else if (regexp.MustCompile(`/quest/user`)).MatchString(r.URL.RequestURI()) {
			h.createUser(w, r)
		}
	case http.MethodGet:
		h.getHistoryOfQuest(w, r)
	case http.MethodPut:
		if (regexp.MustCompile(`/quest/user/donequest/*`)).MatchString(r.URL.RequestURI()) {
			h.changeQuestStatus(w, r)
		} else if (regexp.MustCompile(`/quest/user/donesubquest/*`)).MatchString(r.URL.RequestURI()) {
			h.changeSubquestStatus(w, r)
		} else if (regexp.MustCompile(`/quest/user/updateemail/*`)).MatchString(r.URL.RequestURI()) {
			h.updateEmail(w, r)
		} else if (regexp.MustCompile(`/quest/user/updatepassword/*`)).MatchString(r.URL.RequestURI()) {
			h.updatePassword(w, r)
		} else if (regexp.MustCompile(`/quest/user/updateusername/*`)).MatchString(r.URL.RequestURI()) {
			h.updateUsername(w, r)
		}
	case http.MethodDelete:
		h.deleteUser(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	newUser, err := h.service.UserUseCase.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	log.Printf("createUser is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "user is created "+newUser.Username)
}

func (h *Handler) getHistoryOfQuest(w http.ResponseWriter, r *http.Request) {
	id, err := getUuidFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	history, err := h.service.UserUseCase.GetHistoryOfQuest(id)
	history.User.ID = id

	balance, err := h.service.UserUseCase.ShowBalance(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	history.Balance = balance

	jsonBytes, err := json.Marshal(history)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)
	log.Printf("getHistoryOfQuest is completed")
}

func (h *Handler) attachUser(w http.ResponseWriter, r *http.Request) {

	idQuest, idUser, err := getTwoUuidFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.AttachToQuest(idQuest, idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("attachUser is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user is attached")

}

func (h *Handler) attachUserSubquest(w http.ResponseWriter, r *http.Request) {

	idSubquest, idUser, err := getTwoUuidFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.AttachToSubquest(idSubquest, idUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("attachUserSubquest is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user is attached")
}

func (h *Handler) changeSubquestStatus(w http.ResponseWriter, r *http.Request) {
	idSubquest, idUser, err := getTwoUuidFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.ChangeSubquestsStatus(idSubquest, idUser, status.Done)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("changeSubquestStatus is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user changed status of subquest")
}

func (h *Handler) changeQuestStatus(w http.ResponseWriter, r *http.Request) {
	idQuest, idUser, err := getTwoUuidFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.ChangeQuestsStatus(idQuest, idUser, "done")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("changeQuestStatus is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user changed status of quest")
}

func (h *Handler) updateEmail(w http.ResponseWriter, r *http.Request) {
	idUser, email, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.UpdateEmail(idUser, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateEmail is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user changed email")
}

func (h *Handler) updateUsername(w http.ResponseWriter, r *http.Request) {
	idUser, username, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.UpdateName(idUser, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updateUsername is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user changed username")
}

func (h *Handler) updatePassword(w http.ResponseWriter, r *http.Request) {
	idUser, password, err := getUuidStringFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.UpdatePassword(idUser, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("updatePassword is completed")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "user changed password")
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUuidFromRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	err = h.service.UserUseCase.DeleteUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
	}

	log.Printf("deleteUser is completed")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "user is deleted ")
}
