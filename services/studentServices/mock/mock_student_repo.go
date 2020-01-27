package mock

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormStudentRepository struct{}

func NewGormStudentMockRepo() *GormStudentRepository {
	return &GormStudentRepository{}
}

func (sr *GormStudentRepository) AddStudent(newStudent models.Student) []error {
	return []error{}
}

func (sr *GormStudentRepository) GetStudents() ([]models.Student, []error) {
	students := []models.Student{
		{Id: 1, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 1},
		{Id: 2, FirstName: "nati", MiddleName: "yewondwosen", Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 1},
		{Id: 3, FirstName: "yoni", MiddleName: "endale", Email: "yoni@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 1},
		{Id: 4, FirstName: "moti", MiddleName: "dinsa", Email: "moti@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 1},
		{Id: 5, FirstName: "anwar", MiddleName: "gashaw", Email: "anwar@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 1},
	}
	return students, []error{}
}

func (sr *GormStudentRepository) GetStudentById(id uint) (models.Student, []error) {
	student := models.Student{
		Id: id, Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e",
		ParentId: 1, ClassRoomId: 1, FirstName: "Nathnael",
		SectionId: 1, MiddleName: "Yewondwosen"}
	return student, []error{}
}

func (sr *GormStudentRepository) DeleteStudent(id uint) []error {
	return []error{}
}

func (sr *GormStudentRepository) UpdateStudent(newStudent models.Student) (models.Student, []error) {
	return newStudent, []error{}
}

func (sr *GormStudentRepository) ViewTasks(classRoomId uint, subjectId uint) ([]models.Task, []error) {
	tasks := []models.Task{
		{
			Comments: []models.Comment{
				{Data: "Data1", TaskId: 1, StudentId: 1},
				{Data: "Data1", TaskId: 1, StudentId: 1},
				{Data: "Data1", TaskId: 1, StudentId: 1},
			},
		},
		{
			Comments: []models.Comment{
				{Data: "Data2", TaskId: 2, StudentId: 2},
				{Data: "Data2", TaskId: 2, StudentId: 2},
				{Data: "Data2", TaskId: 2, StudentId: 2},
			},
		},
		{
			Comments: []models.Comment{
				{Data: "Data3", TaskId: 3, StudentId: 3},
				{Data: "Data3", TaskId: 3, StudentId: 3},
				{Data: "Data3", TaskId: 3, StudentId: 3},
			},
		},
		{
			Comments: []models.Comment{
				{Data: "Data4", TaskId: 4, StudentId: 4},
				{Data: "Data4", TaskId: 4, StudentId: 4},
				{Data: "Data4", TaskId: 4, StudentId: 4},
			},
		},
	}
	return tasks, []error{}
}

func (sr *GormStudentRepository) Comment(taskId uint, studentId uint, data string) []error {
	return []error{}
}

func (sr *GormStudentRepository) ViewClass(sectionId uint) ([]models.Student, []error) {
	students := []models.Student{
		{Id: 1, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: sectionId, ClassRoomId: 1, ParentId: 1},
		{Id: 2, FirstName: "nati", MiddleName: "yewondwosen", Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: sectionId, ClassRoomId: 1, ParentId: 1},
		{Id: 3, FirstName: "yoni", MiddleName: "endale", Email: "yoni@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: sectionId, ClassRoomId: 1, ParentId: 1},
		{Id: 4, FirstName: "moti", MiddleName: "dinsa", Email: "moti@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: sectionId, ClassRoomId: 1, ParentId: 1},
		{Id: 5, FirstName: "anwar", MiddleName: "gashaw", Email: "anwar@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: sectionId, ClassRoomId: 1, ParentId: 1},
	}
	return students, []error{}
}

func (sr *GormStudentRepository) ViewResources(subjectId uint) ([]models.Resources, []error) {
	resources := []models.Resources{
		{SubjectId: subjectId, Title: "Resource1", Description: "Resource1", Path: "Resource1"},
		{SubjectId: subjectId, Title: "Resource2", Description: "Resource2", Path: "Resource2"},
		{SubjectId: subjectId, Title: "Resource3", Description: "Resource3", Path: "Resource3"},
	}
	return resources, []error{}
}

func (sr *GormStudentRepository) ViewResult(studentId uint) (models.Student, []error) {
	student := models.Student{}
	result := []models.Result{
		{Id: 1, SubjectId: 1, StudentId: studentId, Assessment: 1, Test: 1, Final: 1, Total: 1},
		{Id: 2, SubjectId: 2, StudentId: studentId, Assessment: 2, Test: 2, Final: 2, Total: 2},
	}
	student.Result = result
	return student, []error{}
}

func (sr *GormStudentRepository) GetHomeRoomTeacher(studentId uint) (models.Teacher, []error) {
	teacher := models.Teacher{Id: 1, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1}
	return teacher, []error{}
}
