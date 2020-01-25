package usecase

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"github.com/nattigy/parentschoolcommunicationsystem/services/studentServices"
)

type StudentUsecase struct {
	studentRepo studentServices.StudentRepository
}

func NewStudentUsecase(StudentRepo studentServices.StudentRepository) *StudentUsecase {
	return &StudentUsecase{studentRepo: StudentRepo}
}

func (sr *StudentUsecase) AddStudent(newStudent models.Student) []error {
	errs := sr.studentRepo.AddStudent(newStudent)
	return errs
}

func (sr *StudentUsecase) GetStudents() ([]models.Student, []error) {
	students, errs := sr.studentRepo.GetStudents()
	return students, errs
}

func (sr *StudentUsecase) GetStudentById(id uint) (models.Student, []error) {
	student, errs := sr.studentRepo.GetStudentById(id)
	return student, errs
}

func (sr *StudentUsecase) DeleteStudent(id uint) []error {
	errs := sr.studentRepo.DeleteStudent(id)
	return errs
}

func (sr *StudentUsecase) UpdateStudent(newStudent models.Student) (models.Student, []error) {
	student, errs := sr.UpdateStudent(newStudent)
	return student, errs
}

func (sr *StudentUsecase) ViewTasks(classRoomId uint, subjectId uint) ([]models.Task, []error) {
	tasks, errs := sr.ViewTasks(classRoomId, subjectId)
	return tasks, errs
}

func (sr *StudentUsecase) Comment(taskId uint, studentId uint, data string) []error {
	errs := sr.studentRepo.Comment(taskId, studentId, data)
	return errs
}

func (sr *StudentUsecase) ViewClass(sectionId uint) ([]models.Student, []error) {
	students, errs := sr.studentRepo.ViewClass(sectionId)
	return students, errs
}

func (sr *StudentUsecase) ViewResources(subjectId uint) ([]models.Resources, []error) {
	resources, errs := sr.studentRepo.ViewResources(subjectId)
	return resources, errs
}

func (sr *StudentUsecase) ViewResult(studentId uint) (models.Student, []error) {
	students, errs := sr.studentRepo.ViewResult(studentId)
	return students, errs
}

func (sr *StudentUsecase) GetHomeRoomTeacher(studentId uint) (models.Teacher, []error) {
	teacher, errs := sr.GetHomeRoomTeacher(studentId)
	return teacher, errs
}
