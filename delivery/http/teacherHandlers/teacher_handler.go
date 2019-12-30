package teacherHandlers

import (
	"github.com/nattigy/parentschoolcommunicationsystem/teacher/usecase"
	"html/template"
	"net/http"
)

type TeacherHandler struct {
	templ    *template.Template
	TUsecase usecase.TeacherUsecase
}

func NewTeacherHandler(t *template.Template, us usecase.TeacherUsecase) *TeacherHandler {
	return &TeacherHandler{
		templ:    t,
		TUsecase: us,
	}
}
func (t *TeacherHandler) MakeNewPost(w http.ResponseWriter, r *http.Request) {

}

func (t *TeacherHandler) EditPost(w http.ResponseWriter, r *http.Request) {

}

func (t *TeacherHandler) RemoveTask(w http.ResponseWriter, r *http.Request) {

}

func (t *TeacherHandler) UploadResource(w http.ResponseWriter, r *http.Request) {

}

func (t *TeacherHandler) TeacherUpdateProfile(w http.ResponseWriter, r *http.Request) {

}

func (t *TeacherHandler) ReportGrade(w http.ResponseWriter, r *http.Request) {

}

func (t *TeacherHandler) ViewClasses(w http.ResponseWriter, r *http.Request) {

}
