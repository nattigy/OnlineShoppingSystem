package teacherHandlers

import (
	"context"
	"encoding/json"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	repository "github.com/nattigy/parentschoolcommunicationsystem/services/session/mock"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session/usecase"
	repository2 "github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices/mock"
	usecase2 "github.com/nattigy/parentschoolcommunicationsystem/services/teacherServices/usecase"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTeacherHandler_AddTeacher(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/admin/teacher/new", nil)

	if err != nil {
		t.Error(err)
	}

	sessionMockRepo := repository.NewSessionMockRepo()
	sessionSer := usecase.NewSessionUsecase(sessionMockRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "teacher", Uuid: "kahdfuhiudhfighiuse", UserID: 1})
	sess, _ := sessionSer.GetSession("kahdfuhiudhfighiuse")
	ctx := context.WithValue(req.Context(), "signed_in_user_session", sess)

	teacherMockRepo := repository2.NewGormTeacherMockRepository()
	teacherServ := usecase2.NewTeacherUsecase(teacherMockRepo)

	templ := template.Must(template.ParseGlob("C:/Users/bek/go/src/github.com/nattigy/parentschoolcommunicationsystem/ui/templates/*.html"))
	thandler := NewTeacherHandler(templ, sessionSer, *teacherServ)

	thandler.AddTeacher(httprr, req.WithContext(ctx))
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	tasks := models.Task{}
	_ = json.Unmarshal([]byte(body), &tasks)

	if string(body) != "About" {
		//
	}
}

func TestTeacherHandler_GetTeachers(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/admin/teachers", nil)

	if err != nil {
		t.Error(err)
	}

	sessionMockRepo := repository.NewSessionMockRepo()
	sessionSer := usecase.NewSessionUsecase(sessionMockRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "teacher", Uuid: "kahdfuhiudhfighiuse", UserID: 1})
	sess, _ := sessionSer.GetSession("kahdfuhiudhfighiuse")
	ctx := context.WithValue(req.Context(), "signed_in_user_session", sess)

	teacherMockRepo := repository2.NewGormTeacherMockRepository()
	teacherServ := usecase2.NewTeacherUsecase(teacherMockRepo)

	templ := template.Must(template.ParseGlob("C:/Users/bek/go/src/github.com/nattigy/parentschoolcommunicationsystem/ui/templates/*.html"))
	thandler := NewTeacherHandler(templ, sessionSer, *teacherServ)

	thandler.GetTeachers(httprr, req.WithContext(ctx))
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	tasks := models.Task{}
	_ = json.Unmarshal([]byte(body), &tasks)

	if string(body) != "About" {
		//
	}
}

func TestTeacherHandler_GetTasks(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/teacher/fetchPosts", nil)

	if err != nil {
		t.Error(err)
	}

	sessionMockRepo := repository.NewSessionMockRepo()
	sessionSer := usecase.NewSessionUsecase(sessionMockRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "teacher", Uuid: "kahdfuhiudhfighiuse", UserID: 1})
	sess, _ := sessionSer.GetSession("kahdfuhiudhfighiuse")
	ctx := context.WithValue(req.Context(), "signed_in_user_session", sess)

	teacherMockRepo := repository2.NewGormTeacherMockRepository()
	teacherServ := usecase2.NewTeacherUsecase(teacherMockRepo)

	templ := template.Must(template.ParseGlob("C:/Users/bek/go/src/github.com/nattigy/parentschoolcommunicationsystem/ui/templates/*.html"))
	thandler := NewTeacherHandler(templ, sessionSer, *teacherServ)

	thandler.GetTasks(httprr, req.WithContext(ctx))
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	tasks := models.Task{}
	_ = json.Unmarshal([]byte(body), &tasks)

	if string(body) != "About" {
		//
	}
}

func TestTeacherHandler_ViewStudents(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/teacher/viewClasses", nil)

	if err != nil {
		t.Error(err)
	}

	sessionMockRepo := repository.NewSessionMockRepo()
	sessionSer := usecase.NewSessionUsecase(sessionMockRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "teacher", Uuid: "kahdfuhiudhfighiuse", UserID: 1})
	sess, _ := sessionSer.GetSession("kahdfuhiudhfighiuse")
	ctx := context.WithValue(req.Context(), "signed_in_user_session", sess)

	teacherMockRepo := repository2.NewGormTeacherMockRepository()
	teacherServ := usecase2.NewTeacherUsecase(teacherMockRepo)

	templ := template.Must(template.ParseGlob("C:/Users/bek/go/src/github.com/nattigy/parentschoolcommunicationsystem/ui/templates/*.html"))
	thandler := NewTeacherHandler(templ, sessionSer, *teacherServ)

	thandler.ViewStudents(httprr, req.WithContext(ctx))
	resp := httprr.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	tasks := models.Task{}
	_ = json.Unmarshal([]byte(body), &tasks)

	if string(body) != "About" {
		//
	}
}
