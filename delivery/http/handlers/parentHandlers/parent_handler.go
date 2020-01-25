package parentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"html/template"
	"net/http"
)

type ParentHandler struct {
	templ   *template.Template
	Session session.SessionUsecase
}

func NewParentHandler(templ *template.Template, session session.SessionUsecase) *ParentHandler {
	return &ParentHandler{templ: templ, Session: session}
}

type ParentInfo struct {
	User   models.User
	Result []models.Result
}

func (ph *ParentHandler) AddParent(w http.ResponseWriter, r *http.Request) {

}

func (ph *ParentHandler) GetParents(w http.ResponseWriter, r *http.Request) {

}

func (ph *ParentHandler) DeleteParent(w http.ResponseWriter, r *http.Request) {

}

func (ph *ParentHandler) ViewGrade(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	in := ParentInfo{
		User:   user,
		Result: []models.Result{},
	}
	err := ph.templ.ExecuteTemplate(w, "parentViewResult", in)
	if err != nil {
		fmt.Println(err)
	}
}
