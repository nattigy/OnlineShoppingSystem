package teacherHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices/usecase"
	"html/template"
	"net/http"
	"strconv"
)

type TeacherHandler struct {
	templ    *template.Template
	Session  session.SessionUsecase
	TUsecase usecase.TeacherUsecase
}

func NewTeacherHandler(templ *template.Template, session session.SessionUsecase, TUsecase usecase.TeacherUsecase) *TeacherHandler {
	return &TeacherHandler{templ: templ, Session: session, TUsecase: TUsecase}
}

type TeacherInfo struct {
	User          models.User
	Resource      models.Resources
	UpdateProfile models.Teacher
	Students      []models.Student
	FetchPost     []models.Task
	Task          models.Task
	Tasks         []models.Task
	Result        []models.Result
	Teacher       models.Teacher
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

	email := r.FormValue("studentEmail")
	password := r.FormValue("studentPassword")
	profile := r.FormValue("studentProfilePic")

	in := TeacherInfo{
		User:          user,
		UpdateProfile: teacher,
	}

	if email != "" || password != "" || profile != "" {
		user.Email = email
		user.Password = password
		teacherUpdateInfo := models.Teacher{Id: user.Id, Email: email, Password: password, ProfilePic: profile}
		newTeaher, errs := th.TUsecase.UpdateTeacher(teacherUpdateInfo)
		if len(errs) > 0 {
			fmt.Println(errs)
		}
		in = TeacherInfo{
			User:          user,
			UpdateProfile: newTeaher,
		}
	}
	err := th.templ.ExecuteTemplate(w, "teacherUpdateProfile.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	user, _ := r.Context().Value("signed_in_user_session").(models.User)
	teacher, _ := th.TUsecase.GetTeacherById(user.Id)
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
		User:    user,
		Teacher: teacher,
	}

	err := th.templ.ExecuteTemplate(w, "teacherMakeNewPost.layout", in)
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
		Tasks:     prevoiusPosts,
	}
	err := th.templ.ExecuteTemplate(w, "teacherEditPost.layout", in)
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
	err := th.templ.ExecuteTemplate(w, "teacherUploadResource.layout", in)
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
	err := th.templ.ExecuteTemplate(w, "teacherReportGrade.layout", in)
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
	err := th.templ.ExecuteTemplate(w, "teacherViewClasses.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}
