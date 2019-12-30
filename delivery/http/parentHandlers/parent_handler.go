package parentHandlers

import (
	"github.com/nattigy/parentschoolcommunicationsystem/parent/usecase"
	"html/template"
	"net/http"
)

type ParentHandler struct {
	templ    *template.Template
	PUsecase usecase.ParentUsecase
}

func NewParentHandler(t *template.Template, us usecase.ParentUsecase) *ParentHandler {
	return &ParentHandler{
		templ:    t,
		PUsecase: us,
	}
}

func (p *ParentHandler) ViewGrade(w http.ResponseWriter, r *http.Request) {

}
