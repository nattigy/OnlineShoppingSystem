package parentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/parent/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/session"
	"github.com/nattigy/parentschoolcommunicationsystem/utility"
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
}
