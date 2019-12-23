package teacherHandlers

import (
	"github.com/nattigy/parentschoolcommunicationsystem/teacher/usecase"
	"net/http"
)

type TeacherHandler struct {
	TUsecase usecase.TeacherUsecase
}

func NewTeacherHandler(h *http.ServeMux, us usecase.TeacherUsecase) {
	handler := &TeacherHandler{
		TUsecase: us,
	}
	h.HandleFunc("/teacher/makeNewPost", handler.MakeNewPost)
	h.HandleFunc("/teacher/editPost", handler.EditPost)
	h.HandleFunc("/teacher/removeTask", handler.RemoveTask)
	h.HandleFunc("/teacher/uploadResources", handler.UploadResource)
	h.HandleFunc("/teacher/updateProfile", handler.TeacherUpdateProfile)
	h.HandleFunc("/teacher/reportGrade", handler.ReportGrade)
	h.HandleFunc("/teacher/viewClasses", handler.ViewClasses)
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
