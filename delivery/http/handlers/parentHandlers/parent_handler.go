package parentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strconv"
)

type ParentHandler struct {
	templ    *template.Template
	Session  session.SessionUsecase
	PUsecase usecase.ParentUsecase
}

func NewParentHandler(templ *template.Template, session session.SessionUsecase, PUsecase usecase.ParentUsecase) *ParentHandler {
	return &ParentHandler{templ: templ, Session: session, PUsecase: PUsecase}
}

type ParentInfo struct {
	User    models.User
	Results []models.Result
	Parents []models.Parent
}

func (ph *ParentHandler) AddParent(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)

	FirstName := r.FormValue("firstname")
	MiddleName := r.FormValue("middlename")
	Email := r.FormValue("email")
	parentid := r.FormValue("parentid")

	if FirstName != "" && MiddleName != "" && Email != "" {
		password, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
		parId, _ := strconv.Atoi(parentid)
		parent := models.Parent{
			Id:         uint(parId),
			FirstName:  FirstName,
			MiddleName: MiddleName,
			Email:      Email,
			Password:   string(password),
		}
		errs := ph.PUsecase.AddParent(parent)
		if errs != nil {
			fmt.Println(errs)
		}
	}
	in := ParentInfo{
		User: models.User{Id: sess.UserID, Email: sess.Email, Role: sess.Role, LoggedIn: true},
	}
	err := ph.templ.ExecuteTemplate(w, "adminAddParent.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (ph *ParentHandler) GetParents(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	parents, errs := ph.PUsecase.GetParents()
	if errs != nil {
		fmt.Println(errs)
	}
	in := ParentInfo{
		User:    models.User{Id: sess.UserID, Email: sess.Email, Role: sess.Role, LoggedIn: true},
		Parents: parents,
	}
	err := ph.templ.ExecuteTemplate(w, "adminListParent.layout", in)
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
	http.Redirect(w, r, "/admin/parents", http.StatusSeeOther)
}

func (ph *ParentHandler) ViewGrade(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	in := ParentInfo{
		User:    models.User{Id: sess.UserID, Role: sess.Role, Email: sess.Email, LoggedIn: true},
		Results: []models.Result{},
	}
	err := ph.templ.ExecuteTemplate(w, "parentViewResult.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}
