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
		{},
		{},
		{},
		{},
		{},
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
				{},
				{},
				{},
			},
		},
		{
			Comments: []models.Comment{
				{},
				{},
				{},
			},
		},
		{
			Comments: []models.Comment{
				{},
				{},
				{},
			},
		},
		{
			Comments: []models.Comment{
				{},
				{},
				{},
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
		{},
		{},
		{},
		{},
	}
	return students, []error{}
}

func (sr *GormStudentRepository) ViewResources(subjectId uint) ([]models.Resources, []error) {
	resources := []models.Resources{
		{},
		{},
		{},
	}
	return resources, []error{}
}

func (sr *GormStudentRepository) ViewResult(studentId uint) (models.Student, []error) {
	student := models.Student{}
	result := []models.Result{
		{},
		{},
	}
	student.Result = result
	return student, []error{}
}

func (sr *GormStudentRepository) GetHomeRoomTeacher(studentId uint) (models.Teacher, []error) {
	teacher := models.Teacher{}
	return teacher, []error{}
}
