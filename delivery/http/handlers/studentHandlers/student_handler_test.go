package studentHandlers

import (
	"fmt"
	"github.com/nattigy/parentschoolcommunicationsystem/database"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session/repository"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session/usecase"
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

	gormdb, err := database.Config()
	if err != nil {
		fmt.Println(err)
	}
	defer gormdb.Close()

	sessionRepo := repository.NewSessionRepository(gormdb)
	sessionSer := usecase.NewSessionUsecase(sessionRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "student", Uuid: "sgdfgf", UserID: 1})

	//user, _ := req.Context().Value("signed_in_user_session").(models.User)
	shandler := NewStudentHandler(templ, sessionSer)

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

	if string(body) != "About" {
		//
	}
}
