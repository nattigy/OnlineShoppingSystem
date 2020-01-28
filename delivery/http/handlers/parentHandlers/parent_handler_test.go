package parentHandlers

import (
	"context"
	"encoding/json"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices/mock"
	usecase2 "github.com/nattigy/parentschoolcommunicationsystem/services/parentServices/usecase"
	repository "github.com/nattigy/parentschoolcommunicationsystem/services/session/mock"
	"github.com/nattigy/parentschoolcommunicationsystem/services/session/usecase"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParentHandler_ViewGrade(t *testing.T) {
	httprr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/parent/viewGrade", nil)

	if err != nil {
		t.Error(err)
	}

	sessionMockRepo := repository.NewSessionMockRepo()
	sessionSer := usecase.NewSessionUsecase(sessionMockRepo)
	sessionSer.CreateSession(httprr, models.Session{Role: "parent", Uuid: "kahdfuhiudhfighiuse", UserID: 1})
	sess, _ := sessionSer.GetSession("kahdfuhiudhfighiuse")
	ctx := context.WithValue(req.Context(), "signed_in_user_session", sess)

	parentMockRepo := mock.NewGormParentMockRepo()
	parentSer := usecase2.NewParentUsecase(parentMockRepo)

	templ := template.Must(template.ParseGlob("C:/Users/bek/go/src/github.com/nattigy/parentschoolcommunicationsystem/ui/templates/*.html"))
	phandler := NewParentHandler(templ, sessionSer, *parentSer)

	phandler.ViewGrade(httprr, req.WithContext(ctx))
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
