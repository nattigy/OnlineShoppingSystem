package parentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"html/template"
	"net/http"
	"strconv"
)

type ParentHandler struct {
	templ    *template.Template
	Session  session.SessionUsecase
	PUsecase parentServices.ParentUsecase
}

func NewParentHandler(templ *template.Template, session session.SessionUsecase) *ParentHandler {
	return &ParentHandler{templ: templ, Session: session}
}

type ParentInfo struct {
	User   models.User
	Result []models.Result
}

func (ph *ParentHandler) AddParent(w http.ResponseWriter, r *http.Request) {
	FirstName := r.FormValue("firstname")
	MiddleName := r.FormValue("middlename")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	ProfilePic := r.FormValue("profilepic")

	if FirstName != "" && MiddleName != "" && Email != "" && Password != "" && ProfilePic != "" {
		parent := models.Parent{
			FirstName:  FirstName,
			MiddleName: MiddleName,
			Email:      Email,
			Password:   Password,
			ProfilePic: ProfilePic,
		}
		errs := ph.PUsecase.AddParent(parent)
		if errs != nil {
			fmt.Println(errs)
		}
	}
	http.Redirect(w, r, "", http.StatusSeeOther)
}

func (ph *ParentHandler) GetParents(w http.ResponseWriter, r *http.Request) {
	parents, errs := ph.PUsecase.GetParents()
	if errs != nil {
		fmt.Println(errs)
	}
	err := ph.templ.ExecuteTemplate(w, "getParents", parents)
	if err != nil {
		fmt.Println(err)
	}
}

func (ph *ParentHandler) DeleteParent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	errs := ph.PUsecase.DeleteParent(uint(id))

	if len(errs) > 0 {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "", http.StatusSeeOther)
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
