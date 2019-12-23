package parentHandlers

import (
	"github.com/nattigy/parentschoolcommunicationsystem/parent/usecase"
	"net/http"
)

type ParentHandler struct {
	PUsecase usecase.ParentUsecase
}

func NewParentHandler(h *http.ServeMux, us usecase.ParentUsecase) {
	handler := &ParentHandler{
		PUsecase: us,
	}
	h.HandleFunc("/teacher/viewGrade", handler.ViewGrade)
}

func (p *ParentHandler) ViewGrade(w http.ResponseWriter, r *http.Request) {

}
