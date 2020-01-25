package teacherHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices"
	"html/template"
	"net/http"
	"strconv"
)

type TeacherHandler struct {
	templ    *template.Template
	Session  session.SessionUsecase
	TUsecase teacherServices.TeacherUsecase
}

func NewTeacherHandler() *TeacherHandler {
	return &TeacherHandler{}
}

type TeacherInfo struct {
	User          models.User
	Resource      models.Resources
	UpdateProfile models.Teacher
	Students      []models.Student
	FetchPost     []models.Task
	Task          models.Task
	Result        []models.Result
}

func (th *TeacherHandler) AddTeacher(w http.ResponseWriter, r *http.Request) {
}

func (th *TeacherHandler) GetTeachers(w http.ResponseWriter, r *http.Request) {
}

func (th *TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
}

func (th *TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	teacher, _ := th.TUsecase.GetTeacherById(user.Id)
	in := TeacherInfo{
		User:          user,
		UpdateProfile: teacher,
	}
	err := th.templ.ExecuteTemplate(w, "teacherUpdateProfile", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)

	title := r.FormValue("title")
	shortDescriptio := r.FormValue("shortDescription")
	description := r.FormValue("description")
	grade, _ := strconv.Atoi(r.FormValue("grade"))
	section := r.FormValue("section")

	teacher, errs := th.TUsecase.GetTeacherById(user.Id)
	newTask := models.Task{Title: title, ShortDescription: shortDescriptio, Description: description, ClassRoomId: teacher.ClassRoomId, SubjectId: teacher.SubjectId}

	if len(errs) > 0 {
		fmt.Println(errs)
	}

	//input validation

	if title != "" && shortDescriptio != "" && description != "" && grade != 0 && section != "" {
		errs = th.TUsecase.CreateTask(newTask)
		if errs != nil {
			fmt.Println(errs)
		}
	}

	in := TeacherInfo{
		User: user,
	}

	err := th.templ.ExecuteTemplate(w, "teacherMakeNewPost", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	teacher, _ := th.TUsecase.GetTeacherById(user.Id)
	prevoiusPosts, errs := th.TUsecase.GetTasks(teacher.SubjectId)
	if errs != nil {
		fmt.Println(errs)
	}
	in := TeacherInfo{
		User:      user,
		FetchPost: prevoiusPosts,
		Task:      prevoiusPosts[0],
	}
	err := th.templ.ExecuteTemplate(w, "teacherEditPost", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	editTitle := r.FormValue("editTitle")
	editDate := r.FormValue("editDate")
	editDescription := r.FormValue("editDescription")
	id, _ := strconv.Atoi(r.FormValue("id"))

	//input validation

	editedTask := models.Task{Id: uint(id), Title: editTitle, Deadline: editDate, ShortDescription: editDescription}

	_, errs := th.TUsecase.UpdateTask(editedTask)
	if errs != nil {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "/teacher/fetchPosts", http.StatusSeeOther)
}

func (th *TeacherHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId, _ := strconv.Atoi(r.FormValue("id"))
	errs := th.TUsecase.DeleteTask(uint(taskId))
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "/teacher/fetchPosts", http.StatusSeeOther)
}

func (th *TeacherHandler) UploadResource(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	in := TeacherInfo{
		User: user,
	}
	err := th.templ.ExecuteTemplate(w, "teacherUploadResource", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) DeleteResource(w http.ResponseWriter, r *http.Request) {
	resourceId, _ := strconv.Atoi(r.FormValue("id"))
	errs := th.TUsecase.DeleteResource(uint(resourceId))
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "/teacher/fetchPosts", http.StatusSeeOther)
}

func (th *TeacherHandler) ReportGrade(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	in := TeacherInfo{
		User: user,
	}
	err := th.templ.ExecuteTemplate(w, "teacherReportGrade", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) ViewStudents(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	students, errs := th.TUsecase.ViewStudents(user.Id)
	if errs != nil {
		fmt.Println(errs)
	}
	in := TeacherInfo{
		User:     user,
		Students: students,
	}
	err := th.templ.ExecuteTemplate(w, "teacherViewClasses", in)
	if err != nil {
		fmt.Println(err)
	}
}
