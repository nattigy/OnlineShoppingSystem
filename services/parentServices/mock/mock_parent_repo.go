package mock

import "github.com/nattigy/parentschoolcommunicationsystem/models"

type GormParentRepository struct{}

func NewGormParentMockRepo() *GormParentRepository {
	return &GormParentRepository{}
}
func (pr *GormParentRepository) AddParent(parent models.Parent) []error {
	return []error{}
}

func (pr *GormParentRepository) GetParents() ([]models.Parent, []error) {
	parents := []models.Parent{
		{Id: 1, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e"},
		{Id: 1, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e"},
		{Id: 1, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e"},
	}
	return parents, []error{}
}

func (pr *GormParentRepository) GetParentById(id uint) (models.Parent, []error) {
	parents := models.Parent{
		Id: id, FirstName: "bek", MiddleName: "zemed",
		Email:    "bek@gmail.com",
		Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e",
	}
	return parents, []error{}
}

func (pr *GormParentRepository) DeleteParent(id uint) []error {
	return []error{}
}

func (pr *GormParentRepository) UpdateParent(newParent models.Parent) (models.Parent, []error) {
	return newParent, []error{}
}

func (pr *GormParentRepository) GetChild(parentId uint) (models.Student, []error) {
	student := models.Student{Id: 2, FirstName: "bek", MiddleName: "zemed", Email: "bek@gmail.com", Password: "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e", SectionId: 1, ClassRoomId: 1, ParentId: parentId}
	return student, []error{}
}

func (pr *GormParentRepository) ViewGrade(studentId uint) (models.Student, []error) {
	grade := models.Student{
		Id: 2, FirstName: "bek",
		MiddleName:  "zemed",
		Email:       "bek@gmail.com",
		Password:    "$2a$10$izeCetsu3s9pBSJmRDlfzeXCpblROeKhVwUMpruzCIpUDob3QbI.e",
		SectionId:   1,
		ClassRoomId: 1,
		ParentId:    1,
		Result: []models.Result{
			{StudentId: studentId, Assessment: 1, Test: 1, Final: 1, Total: 1},
			{StudentId: studentId, Assessment: 1, Test: 1, Final: 1, Total: 1},
		},
	}
	return grade, []error{}

}
