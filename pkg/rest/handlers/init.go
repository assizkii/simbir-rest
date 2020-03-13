package handlers

import (
	"encoding/json"
	"github.com/assizkii/simbir-rest/internal/domain/interfaces"
	"github.com/astaxie/beego/session"
	"net/http"
)

type RestHandler struct {
	store interfaces.AppStorage
}

var globalSessions *session.Manager

func Init(st *interfaces.AppStorage) *RestHandler {
	return &RestHandler{store: *st}
}

type HttpResponse struct {
	Status int
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func showResponse(result *HttpResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(result.Status)
	w.Write(response)
}

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:              "gosessionid",
		EnableSetCookie:         true,
		Gclifetime:              3600,
		Maxlifetime:             3600,
		DisableHTTPOnly:         false,
		Secure:                  false,
		CookieLifeTime:          3600,
		ProviderConfig:          "",
		Domain:                  "",
		SessionIDLength:         0,
		EnableSidInHTTPHeader:   false,
		SessionNameInHTTPHeader: "",
		EnableSidInURLQuery:     false,
		SessionIDPrefix:         "",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func checkSession(w http.ResponseWriter, r *http.Request) bool {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	login := sess.Get("login")
	return login != nil
}

func setSession(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	sess.Set("login", r.Form["login"])
}
