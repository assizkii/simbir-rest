package handlers

import (
	"github.com/assizkii/simbir-rest/internal/domain/usecases"
	"github.com/assizkii/simbir-rest/internal/entities"
	"net/http"
)

func (handler *RestHandler) Registration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user entities.User
	user.Login = r.FormValue("login")
	user.Password = r.FormValue("password")

	err := handler.store.Add(user)
	if err != nil {
		result := &HttpResponse{http.StatusBadRequest, "", err.Error()}
		showResponse(result, w)
		return
	}

	result := &HttpResponse{http.StatusOK, "user registration success", ""}
	showResponse(result, w)
}

func (handler *RestHandler) Auth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.FormValue("login")
	password := r.FormValue("password")

	user, err := handler.store.Get(login)
	if err != nil {
		result := &HttpResponse{http.StatusBadRequest, "", err.Error()}
		showResponse(result, w)
		return
	}

	if !usecases.CheckPassword(password, user.Password) {
		result := &HttpResponse{http.StatusBadRequest, "", "Incorrect password"}
		showResponse(result, w)
		return
	}
	setSession(w, r)
	result := &HttpResponse{http.StatusOK, "user auth success", ""}
	showResponse(result, w)
}

func (handler *RestHandler) Logout(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionDestroy(w, r)
	result := &HttpResponse{http.StatusOK, "logout success", ""}
	showResponse(result, w)
}

func (handler *RestHandler) GetRandNumber(w http.ResponseWriter, r *http.Request) {
	if !checkSession(w, r) {
		result := &HttpResponse{http.StatusBadRequest, "", "Need auth"}
		showResponse(result, w)
		return
	}

	randNumber := usecases.GetRandNumber()

	result := &HttpResponse{http.StatusOK, randNumber, ""}
	showResponse(result, w)

}
