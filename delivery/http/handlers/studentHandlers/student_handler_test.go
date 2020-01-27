package studentHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	repository2 "github.com/nattigy/parentschoolcommunicationsystem/services/session/mock"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session/usecase"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices/mock"
	usecase2 "github.com/nattigy/parentschoolcommunicationsystem/services/studentServices/usecase"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var templ = template.Must(template.ParseGlob("ui/templates/*.html"))

func TestStudentHandler_ViewTasks(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/student/viewTask", nil)

	if err != nil {
		t.Error(err)
	}

	sessionMockRepo := repository2.NewSessionMockRepo()
	sessionSer := usecase.NewSessionUsecase(sessionMockRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "student", Uuid: "sgdfgf", UserID: 1})

	studentMockRepo := mock.NewGormStudentMockRepo()
	studentSer := usecase2.NewStudentUsecase(studentMockRepo)

	//user, _ := req.Context().Value("signed_in_user_session").(models.User)
	shandler := NewStudentHandler(templ, sessionSer, *studentSer)

	shandler.ViewTasks(httprr, req)
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

	fmt.Println(tasks)
	if string(body) != "About" {
		//
	}
}
