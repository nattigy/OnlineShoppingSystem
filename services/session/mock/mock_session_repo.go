package repository

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type SessionRepository struct{}

func NewSessionMockRepo() *SessionRepository {
	return &SessionRepository{}
}

func (s *SessionRepository) CreateSession(session models.Session) (models.Session, []error) {
	return session, []error{}
}

func (s *SessionRepository) GetSession(value string) (models.Session, []error) {
	sess := models.Session{Role: "student", Email: "nati@gmail.com", UserID: 1, Uuid: value}
	return sess, []error{}
}

func (s *SessionRepository) Sessions() ([]models.Session, []error) {
	sess := []models.Session{
		{Role: "student", Email: "nati@gmail.com", UserID: 1, Uuid: "k;ahdfuhiudhfighiuse"},
		{Role: "teacher", Email: "aman@gmail.com", UserID: 1, Uuid: "k;ahdfuhiudhfighiuse"},
	}
	return sess, []error{}
}

func (s *SessionRepository) DeleteSession(id uint) []error {
	return []error{}
}

func (s *SessionRepository) UpdateSession(sess models.Session) (models.Session, []error) {
	return sess, []error{}
}

func (s *SessionRepository) GetUser(id uint) (models.User, []error) {
	user := models.User{Id: id, Email: "nati@gmail.com", Role: "student", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", LoggedIn: true}
	return user, []error{}
}
