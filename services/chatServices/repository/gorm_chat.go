package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type ChatRepository struct {
	conn *gorm.DB
}

func NewChatRepository(conn *gorm.DB) *ChatRepository {
	return &ChatRepository{conn: conn}
}

func (cr *ChatRepository) Store(parentId uint, teacherId uint, data string, sender string) []error {
	newMessage := models.Message{ParentId: parentId, TeacherId: teacherId, MessageContent: data, From: sender}
	errs := cr.conn.Create(&newMessage).GetErrors()
	return errs
}

func (cr *ChatRepository) Get(parentId uint, teacherId uint) ([]models.Message, []error) {
	var messages []models.Message
	errs := cr.conn.Where("parent_id = ? AND teacher_id = ?", parentId, teacherId).Order("created_at").Find(&messages).GetErrors()
	return messages, errs
}
