package teacherHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/validateInput"
	"golang.org/x/crypto/bcrypt"
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

type Error struct {
	Title            string
	Description      string
	ShortDescription string
}

type TeacherInfo struct {
	User          models.User
	Session       models.Session
	Resource      models.Resources
	Resources     []models.Resources
	UpdateProfile models.Teacher
	Students      []models.Student
	FetchPost     []models.Task
	Task          models.Task
	Tasks         []models.Task
	Result        []models.Result
	Teacher       models.Teacher
	Error         Error
	Teachers      []models.Teacher
}

func (th *TeacherHandler) AddTeacher(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	FirstName := r.FormValue("firstname")
	teacherId := r.FormValue("teacherid")
	MiddleName := r.FormValue("middlename")
	Email := r.FormValue("email")
	SubjectId := r.FormValue("subjectid")
	ClassRoomId := r.FormValue("classroomid")

	if FirstName != "" && MiddleName != "" && Email != "" && SubjectId != "" && ClassRoomId != "" {
		id, _ := strconv.Atoi(teacherId)
		password, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
		subject, _ := strconv.Atoi(SubjectId)
		newSubject := uint(subject)
		classRoom, _ := strconv.Atoi(ClassRoomId)
		newClassRoom := uint(classRoom)
		teacher := models.Teacher{
			Id:          uint(id),
			FirstName:   FirstName,
			MiddleName:  MiddleName,
			Email:       Email,
			Password:    string(password),
			SubjectId:   newSubject,
			ClassRoomId: newClassRoom,
		}
		errs := th.TUsecase.AddTeacher(teacher)
		if errs != nil {
			fmt.Println(errs)
		}
	}
	in := TeacherInfo{
		User: models.User{Role: sess.Role, Id: sess.UserID, Email: sess.Email, LoggedIn: true},
	}
	err := th.templ.ExecuteTemplate(w, "adminAddTeacher.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) GetTeachers(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	teachers, errs := th.TUsecase.GetTeachers()
	if errs != nil {
		fmt.Println(errs)
	}
	in := TeacherInfo{
		User:     models.User{Role: sess.Role, Id: sess.UserID, Email: sess.Email, LoggedIn: true},
		Teachers: teachers,
	}
	err := th.templ.ExecuteTemplate(w, "adminListTeacher.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	errs := th.TUsecase.DeleteTeacher(uint(id))

	if len(errs) > 0 {
		fmt.Println(errs)
	}
	http.Redirect(w, r, "/admin/teachers", http.StatusSeeOther)
}

func (th *TeacherHandler) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	teacher, _ := th.TUsecase.GetTeacherById(sess.UserID)
	var user models.User
	email := r.FormValue("studentEmail")
	password := r.FormValue("studentPassword")
	profile := r.FormValue("studentProfilePic")

	user.Email = sess.Email
	user.Role = sess.Role
	user.LoggedIn = true

	in := TeacherInfo{
		User:          user,
		UpdateProfile: teacher,
		Session:       sess,
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
			Session:       sess,
		}
	}
	err := th.templ.ExecuteTemplate(w, "teacherUpdateProfile.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	teacher, _ := th.TUsecase.GetTeacherById(sess.UserID)

	title := r.FormValue("title")
	shortDescription := r.FormValue("shortDescription")
	description := r.FormValue("description")
	grade, _ := strconv.Atoi(r.FormValue("grade"))
	section := r.FormValue("section")
	token := r.FormValue("csrf")

	taskValidation := validateInput.Input{VErrors: validateInput.ValidationErrors{}}
	taskValidation.MatchesPattern(title, validateInput.StringRX)
	taskValidation.MatchesPattern(shortDescription, validateInput.StringRX)
	taskValidation.MatchesPattern(description, validateInput.StringRX)
	taskValidation.MatchesPattern(section, validateInput.StringRX)
	taskValidation.CSRFCheck(token, sess)

	teacher, errs := th.TUsecase.GetTeacherById(sess.UserID)
	newTask := models.Task{Title: title, ShortDescription: shortDescription, Description: description, ClassRoomId: teacher.ClassRoomId, SubjectId: teacher.SubjectId}

	if len(errs) > 0 {
		fmt.Println(errs)
	}

	if len(taskValidation.VErrors) > 0 {
		tt := taskValidation.VErrors[title][0]
		ss := taskValidation.VErrors[shortDescription][0]
		dd := taskValidation.VErrors[description][0]
		csrf := taskValidation.VErrors[token][0]
		fmt.Println(csrf)
		in := TeacherInfo{
			User:    models.User{Role: sess.Role, Email: sess.Email, Id: sess.UserID, LoggedIn: true},
			Teacher: teacher,
			Session: sess,
			Error:   Error{Title: tt, ShortDescription: ss, Description: dd},
		}
		_ = th.templ.ExecuteTemplate(w, "teacherMakeNewPost.layout", in)
	}

	if len(taskValidation.VErrors) == 0 && title != "" && shortDescription != "" && description != "" && grade != 0 && section != "" {
		errs = th.TUsecase.CreateTask(newTask)
		if errs != nil {
			fmt.Println(errs)
		}
	}

	in := TeacherInfo{
		User:    models.User{Role: sess.Role, Email: sess.Email, Id: sess.UserID, LoggedIn: true},
		Teacher: teacher,
		Session: sess,
		Error:   Error{},
	}

	err := th.templ.ExecuteTemplate(w, "teacherMakeNewPost.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	teacher, _ := th.TUsecase.GetTeacherById(sess.UserID)
	prevoiusPosts, errs := th.TUsecase.GetTasks(teacher.SubjectId)
	if errs != nil {
		fmt.Println(errs)
	}
	in := TeacherInfo{
		User:      models.User{Email: sess.Email, Id: sess.UserID, Role: sess.Role, LoggedIn: true},
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
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	user := models.User{Email: sess.Email, Id: sess.UserID, Role: sess.Role, LoggedIn: true}
	teacher, _ := th.TUsecase.GetTeacherById(sess.UserID)
	resources, _ := th.TUsecase.GetResource(teacher.SubjectId)

	//title := r.FormValue("title")
	//description := r.FormValue("description")
	//link := r.FormValue("link")

	if title != "" && description != "" && link != "" {
		err := th.TUsecase.UploadResource(models.Resources{SubjectId: teacher.SubjectId, Title: title, Description: description, Link: link})
		if len(err) > 0 {
			fmt.Println(err)
		}
	}

	in := TeacherInfo{
		User:      user,
		Resources: resources,
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
	http.Redirect(w, r, "/teacher/uploadResources", http.StatusSeeOther)
}

func (th *TeacherHandler) ReportGrade(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	teacher, _ := th.TUsecase.GetTeacherById(sess.UserID)
	user := models.User{Email: sess.Email, Id: sess.UserID, Role: sess.Role, LoggedIn: true}
	in := TeacherInfo{
		User: user,
	}
	fmt.Println("in here")
	studentId := r.FormValue("studentid")
	assement := r.FormValue("assesment")
	test := r.FormValue("test")
	final := r.FormValue("final")
	total := r.FormValue("total")

	if studentId != "" && test != "" && assement != "" && final != "" && total != "" {
		stdId, _ := strconv.Atoi(studentId)
		ass, _ := strconv.Atoi(assement)
		tes, _ := strconv.Atoi(test)
		fin, _ := strconv.Atoi(final)
		tot, _ := strconv.Atoi(total)
		fmt.Println(stdId, tot, ass, fin, tes)
		errs := th.TUsecase.ReportGrade(models.Result{StudentId: uint(stdId), SubjectId: teacher.SubjectId, Total: tot, Final: fin, Test: tes, Assessment: ass})
		if errs != nil {
			fmt.Println(errs)
		}
	}
	err := th.templ.ExecuteTemplate(w, "teacherReportGrade.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (th *TeacherHandler) ViewStudents(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	user := models.User{Email: sess.Email, Id: sess.UserID, Role: sess.Role, LoggedIn: true}
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
