package parentApi

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nattigy/parentschoolcommunicationsystem/services/parentServices"
	"net/http"
)

type ParentApi struct {
	parentService parentServices.ParentUsecase
}

func NewParentApi(parentService parentServices.ParentUsecase) *ParentApi {
	return &ParentApi{parentService: parentService}
}

func (pa *ParentApi) ViewGrade(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
