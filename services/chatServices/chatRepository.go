package chatServices

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type ChatRepository interface {
	Store(parentId uint, teacherId uint, data string, sender string) []error
	Get(parentId uint, teacherId uint) ([]models.Message, []error)
}
