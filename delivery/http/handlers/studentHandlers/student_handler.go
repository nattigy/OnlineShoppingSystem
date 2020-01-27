package studentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices/usecase"
	"html/template"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	templ    *template.Template
	Session  session.SessionUsecase
	SUsecase usecase.StudentUsecase
}

func NewStudentHandler(templ *template.Template, session session.SessionUsecase, SUsecase usecase.StudentUsecase) *StudentHandler {
	return &StudentHandler{templ: templ, Session: session, SUsecase: SUsecase}
}

type StudentInfo struct {
	User          models.User
	Student       models.Student
	Students      []models.Student
	Tasks         []models.Task
	Task          models.Task
	UpdateProfile models.Student
	Resources     []models.Resources
	Result        []models.Result
	ClassMates    []models.Student
}

func (sh *StudentHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	firstName := r.FormValue("firstname")
	middleName := r.FormValue("middlename")
	email := r.FormValue("email")
	password := r.FormValue("password")
	sectionId := r.FormValue("section")
	classRoomId := r.FormValue("classroomid")

	if firstName != "" && middleName != "" && email != "" && password != "" && sectionId != "" && classRoomId != "" {
		secID, _ := strconv.Atoi(sectionId)
		classId, _ := strconv.Atoi(classRoomId)
		student := models.Student{FirstName: firstName, MiddleName: middleName, Email: email, Password: password, SectionId: uint(secID), ClassRoomId: uint(classId)}
		err := sh.SUsecase.AddStudent(student)
		if len(err) > 0 {
			fmt.Println(err)
		}
	}
	in := StudentInfo{
		User: models.User{Id: sess.UserID, Role: sess.Role, LoggedIn: true},
	}
	errr := sh.templ.ExecuteTemplate(w, "adminAddStudent.layout", in)
	if errr != nil {
		fmt.Println(errr)
	}
}

func (sh *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	students, err := sh.SUsecase.GetStudents()
	if len(err) > 0 {
		fmt.Println(err)
	}
	in := StudentInfo{
		User:     models.User{Id: sess.UserID, Role: sess.Role, LoggedIn: true},
		Students: students,
	}
	errr := sh.templ.ExecuteTemplate(w, "adminListStudent.layout", in)
	if errr != nil {
		fmt.Println(errr)
	}

}

func (sh *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("studentid")
	if id != "" {
		converted, _ := strconv.Atoi(id)
		err := sh.SUsecase.DeleteStudent(uint(converted))
		if err != nil {
			fmt.Println(err)
		}
	}
	http.Redirect(w, r, "/admin/students", http.StatusSeeOther)
}

func (sh *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	student, _ := sh.SUsecase.GetStudentById(sess.UserID)
	user := models.User{Role: sess.Role, Email: student.Email, LoggedIn: true}
	email := r.FormValue("studentEmail")
	password := r.FormValue("studentPassword")
	profile := r.FormValue("studentProfilePic")

	in := StudentInfo{
		Student:       student,
		User:          user,
		UpdateProfile: models.Student{Email: student.Email, Password: student.Password},
	}

	if email != "" || password != "" || profile != "" {
		student.Email = email
		student.Password = password
		studentUpdateInfo := models.Student{Id: student.Id, Email: email, Password: password, ProfilePic: profile}
		newStudent, errs := sh.SUsecase.UpdateStudent(studentUpdateInfo)
		if len(errs) > 0 {
			fmt.Println(errs)
		}
		in = StudentInfo{
			User:          user,
			UpdateProfile: newStudent,
		}
	}
	err := sh.templ.ExecuteTemplate(w, "studentUpdateProfile.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) ViewTasks(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	subjectId, _ := strconv.Atoi(r.FormValue("subjectId"))
	//input validation
	student, errs := sh.SUsecase.GetStudentById(sess.UserID)
	fmt.Println(sess.Role)
	user := models.User{Id: student.Id, Email: student.Email, Role: sess.Role, LoggedIn: true}
	data, _ := sh.SUsecase.ViewTasks(student.ClassRoomId, uint(subjectId))
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	in := StudentInfo{
		Tasks: data,
		Task:  models.Task{},
		User:  user,
	}
	err := sh.templ.ExecuteTemplate(w, "studentViewTask.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) Comment(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	student, _ := sh.SUsecase.GetStudentById(sess.UserID)
	comment := r.FormValue("comment")
	taskId, _ := strconv.Atoi(r.FormValue("taskId"))
	//input validation
	_ = sh.SUsecase.Comment(uint(taskId), sess.UserID, student.FirstName, comment)
	http.Redirect(w, r, "/student/viewTask?subjectId=1", http.StatusSeeOther)
}

func (sh *StudentHandler) ViewClass(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	student, _ := sh.SUsecase.GetStudentById(sess.UserID)
	user := models.User{Email: student.Email, Role: sess.Role, Id: student.Id, LoggedIn: true}
	classMates, _ := sh.SUsecase.ViewClass(student.ClassRoomId)
	in := StudentInfo{
		User:       user,
		ClassMates: classMates,
	}
	err := sh.templ.ExecuteTemplate(w, "studentClassMates.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) ViewResources(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	subjectId, _ := strconv.Atoi(r.FormValue("subjectId"))
	resources, errs := sh.SUsecase.ViewResources(uint(subjectId))
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	in := StudentInfo{
		User:      models.User{Id: sess.ID, Role: sess.Role, Email: sess.Email},
		Resources: resources,
	}
	err := sh.templ.ExecuteTemplate(w, "studentResources.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}

func (sh *StudentHandler) ViewResult(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("signed_in_user_session").(models.Session)
	results, _ := sh.SUsecase.ViewResult(sess.UserID)
	in := StudentInfo{
		User:   models.User{Role: sess.Role, Email: sess.Email, Id: sess.UserID},
		Result: results.Result,
	}
	err := sh.templ.ExecuteTemplate(w, "studentViewResult.layout", in)
	if err != nil {
		fmt.Println(err)
	}
}
