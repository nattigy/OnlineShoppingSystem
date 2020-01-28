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
		{Id: 2, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 1},
		{Id: 1, FirstName: "nati", MiddleName: "yewondwosen", Email: "nati@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 2},
		{Id: 3, FirstName: "yoni", MiddleName: "endale", Email: "yoni@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 3},
		{Id: 4, FirstName: "moti", MiddleName: "dinsa", Email: "moti@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 4},
		{Id: 5, FirstName: "anwar", MiddleName: "gashaw", Email: "anwar@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: 5},
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
			Id:               1,
			ClassRoomId:      classRoomId,
			SubjectId:        subjectId,
			Description:      "some text here",
			ShortDescription: "some text here",
			Title:            "task 1",
			Comments: []models.Comment{
				{Data: "comment data 1", TaskId: 1, StudentId: 1},
				{Data: "comment data 1", TaskId: 1, StudentId: 1},
				{Data: "comment data 1", TaskId: 1, StudentId: 1},
			},
		},
		{
			Id:               2,
			ClassRoomId:      classRoomId,
			SubjectId:        subjectId,
			Description:      "some text here",
			ShortDescription: "some text here",
			Title:            "task 2",
			Comments: []models.Comment{
				{Data: "comment data 2", TaskId: 2, StudentId: 2},
				{Data: "comment data 2", TaskId: 2, StudentId: 2},
				{Data: "comment data 2", TaskId: 2, StudentId: 2},
			},
		},
		{
			Id:               3,
			ClassRoomId:      classRoomId,
			SubjectId:        subjectId,
			Description:      "some text here",
			ShortDescription: "some text here",
			Title:            "task 3",
			Comments: []models.Comment{
				{Data: "comment data 3", TaskId: 3, StudentId: 3},
				{Data: "comment data 3", TaskId: 3, StudentId: 3},
				{Data: "comment data 3", TaskId: 3, StudentId: 3},
			},
		},
		{
			Id:               4,
			ClassRoomId:      classRoomId,
			SubjectId:        subjectId,
			Description:      "some text here",
			ShortDescription: "some text here",
			Title:            "task 4",
			Comments: []models.Comment{
				{Data: "comment data 4", TaskId: 4, StudentId: 4},
				{Data: "comment data 4", TaskId: 4, StudentId: 4},
				{Data: "comment data 4", TaskId: 4, StudentId: 4},
			},
		},
	}
	return tasks, []error{}
}

func (sr *GormStudentRepository) Comment(taskId uint, studentId uint, studentName string, data string) []error {
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
		{SubjectId: subjectId, Title: "Resource1", Description: "Resource1", Link: "Resource1"},
		{SubjectId: subjectId, Title: "Resource2", Description: "Resource2", Link: "Resource2"},
		{SubjectId: subjectId, Title: "Resource3", Description: "Resource3", Link: "Resource3"},
	}
	return resources, []error{}
}

func (sr *GormStudentRepository) ViewResult(studentId uint) (models.Student, []error) {
	student := models.Student{Id: 1, ClassRoomId: 1, MiddleName: "Yewondwosen", SectionId: 1, FirstName: "Nathnael", ParentId: 2, Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", Email: "nati@gmail.com"}
	result := []models.Result{
		{Id: 1, SubjectId: 100, StudentId: studentId, Assessment: 19, Test: 29, Final: 40, Total: 88},
		{Id: 2, SubjectId: 101, StudentId: studentId, Assessment: 16, Test: 25, Final: 39, Total: 80},
	}
	student.Result = result
	return student, []error{}
}

func (sr *GormStudentRepository) GetHomeRoomTeacher(studentId uint) (models.Teacher, []error) {
	teacher := models.Teacher{Id: 1, FirstName: "Amanuel", MiddleName: "kebede", Email: "aman@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SubjectId: 1, ClassRoomId: 1}
	return teacher, []error{}
}

func (sr *GormStudentRepository) Comments(taskId uint) ([]models.Comment, []error) {
	comments := []models.Comment{
		{Data: "comment data 4", TaskId: 4, StudentId: 4},
		{Data: "comment data 4", TaskId: 4, StudentId: 4},
		{Data: "comment data 4", TaskId: 4, StudentId: 4},
	}
	return comments, []error{}
}
