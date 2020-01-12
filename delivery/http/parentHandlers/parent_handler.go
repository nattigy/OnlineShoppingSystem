package parentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parent/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/utility"
	"html/template"
	"net/http"
)

type ParentHandler struct {
	templ    *template.Template
	PUsecase usecase.ParentUsecase
	Session  session.SessionUsecase
	utility  utility.Utility
}

func NewParentHandler(templ *template.Template, PUsecase usecase.ParentUsecase, session session.SessionUsecase, utility utility.Utility) *ParentHandler {
	return &ParentHandler{templ: templ, PUsecase: PUsecase, Session: session, utility: utility}
}

type ParentInfo struct {
	User   models.User
	Result []models.Result
}

func (p *ParentHandler) ViewGrade(w http.ResponseWriter, r *http.Request) {
	user, err := p.Session.Check(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	if user.Id == 0 {
		fmt.Println("Id not found")
		return
	}

	in := ParentInfo{
		User:   user,
		Result: []models.Result{},
	}
	//_ = json.NewEncoder(w).Encode(data)
	err = p.templ.ExecuteTemplate(w, "parentViewResult", in)
	if err != nil {
		fmt.Println(err)
	}
}
