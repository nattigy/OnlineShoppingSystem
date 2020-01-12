package usecases

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type ChatUsecase struct {
	conn *gorm.DB
}

func NewChatUsecase(conn *gorm.DB) *ChatUsecase {
	return &ChatUsecase{conn: conn}
}

func (c *ChatUsecase) Store(parent models.Parent, teacher models.Teacher, data string) []error {
	message := models.Message{ParentId: parent.Id, TeacherId: teacher.Id, MessageContent: data}
	errs := c.conn.Create(message).GetErrors()
	return errs
}

func (c *ChatUsecase) Get(parent models.Parent, teacher models.Teacher) ([]models.Message, []error) {
	var messages []models.Message
	errs := c.conn.Where("parent_id = ? AND teacher_id = ?", parent.Id, teacher.Id).Order("created_at").Find(&messages).GetErrors()
	return messages, errs
}
