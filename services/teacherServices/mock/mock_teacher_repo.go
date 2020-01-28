package repository

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormTeacherRepository struct{}

func NewGormTeacherMockRepository() *GormTeacherRepository {
	return &GormTeacherRepository{}
}

func (tr *GormTeacherRepository) AddTeacher(newTeacher models.Teacher) []error {
	return []error{}
}

func (tr *GormTeacherRepository) GetTeachers() ([]models.Teacher, []error) {
	teachers := []models.Teacher{
		{Id: 2, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1},
		{Id: 1, FirstName: "nati", MiddleName: "yewondwosen", Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1},
		{Id: 3, FirstName: "yoni", MiddleName: "endale", Email: "yoni@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1},
		{Id: 4, FirstName: "moti", MiddleName: "dinsa", Email: "moti@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1},
		{Id: 5, FirstName: "anwar", MiddleName: "gashaw", Email: "anwar@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1},
	}
	return teachers, []error{}

}

func (tr *GormTeacherRepository) GetTeacherById(id uint) (models.Teacher, []error) {
	teacher := models.Teacher{
		Id: id, Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e",
		SubjectId: 1, ClassRoomId: 1, FirstName: "Nathnael",
		MiddleName: "Yewondwosen"}
	return teacher, []error{}
}

func (tr *GormTeacherRepository) DeleteTeacher(id uint) []error {
	return []error{}
}

func (tr *GormTeacherRepository) UpdateTeacher(newTeacher models.Teacher) (models.Teacher, []error) {
	return newTeacher, []error{}
}

func (tr *GormTeacherRepository) CreateTask(task models.Task) []error {
	return []error{}
}

func (tr *GormTeacherRepository) GetTasks(subjectId uint) ([]models.Task, []error) {
	tasks := []models.Task{
		{Id: 1, Title: "GO", Description: "go is the best programming language", ShortDescription: "go is the best", SubjectId: subjectId, ClassRoomId: 1, Deadline: "tommorow"},
		{Id: 1, Title: "GO", Description: "go is the best programming language", ShortDescription: "go is the best", SubjectId: subjectId, ClassRoomId: 1, Deadline: "tommorow"},
	}
	return tasks, []error{}
}

func (tr *GormTeacherRepository) UpdateTask(newTask models.Task) (models.Task, []error) {
	return newTask, []error{}
}

func (tr *GormTeacherRepository) DeleteTask(taskId uint) []error {
	return []error{}
}

func (tr *GormTeacherRepository) UploadResource(resource models.Resources) []error {
	return []error{}
}

func (tr *GormTeacherRepository) DeleteResource(resourceId uint) []error {
	return []error{}
}

func (tr *GormTeacherRepository) ReportGrade(grade models.Result) []error {
	return []error{}
}

func (tr *GormTeacherRepository) ViewStudents(classRoomId uint) ([]models.Student, []error) {
	students := []models.Student{
		{Id: 2, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: classRoomId},
		{Id: 1, FirstName: "nati", MiddleName: "yewondwosen", Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: classRoomId},
		{Id: 3, FirstName: "yoni", MiddleName: "endale", Email: "yoni@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: classRoomId},
		{Id: 4, FirstName: "moti", MiddleName: "dinsa", Email: "moti@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: classRoomId},
		{Id: 5, FirstName: "anwar", MiddleName: "gashaw", Email: "anwar@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: classRoomId},
	}
	return students, []error{}

}

func (tr *GormTeacherRepository) GetResource(subjectId uint) ([]models.Resources, []error) {
	resource := []models.Resources{
		{SubjectId: subjectId, Title: "GO", Description: "go is the best programming language", Link: "github.com/pcscs"},
		{SubjectId: subjectId, Title: "GO", Description: "go is the best programming language", Link: "github.com/pcscs"},
		{SubjectId: subjectId, Title: "GO", Description: "go is the best programming language", Link: "github.com/pcscs"},
	}
	return resource, []error{}
}
