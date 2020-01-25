package usecases

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/chatServices"
)

type ChatUsecase struct {
	chatRepo chatServices.ChatRepository
}

func NewChatUsecase(ChatRepo chatServices.ChatRepository) *ChatUsecase {
	return &ChatUsecase{chatRepo: ChatRepo}
}

func (c *ChatUsecase) Store(parentId uint, teacherId uint, data string, sender string) []error {
	errs := c.chatRepo.Store(parentId, teacherId, data, sender)
	return errs
}

func (c *ChatUsecase) Get(parentId uint, teacherId uint) ([]models.Message, []error) {
	messages, errs := c.chatRepo.Get(parentId, teacherId)
	return messages, errs
}
