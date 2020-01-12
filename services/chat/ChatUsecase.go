package chat

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type ChatUsecase interface {
	Store(parent models.Parent, teacher models.Teacher, data string) []error
	Get(parent models.Parent, teacher models.Teacher) ([]models.Message, []error)
}
