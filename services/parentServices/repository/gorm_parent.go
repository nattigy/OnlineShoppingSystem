package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nattigy/parentschoolcommunicationsystem/models"
)

type GormParentRepository struct {
	conn *gorm.DB
}

func NewGormParentRepository(Conn *gorm.DB) *GormParentRepository {
	return &GormParentRepository{conn: Conn}
}

func (pr *GormParentRepository) AddParent(parent models.Parent) []error {
	user := models.User{Id: parent.Id, Role: "parent", Email: parent.Email, Password: parent.Password}
	errs := pr.conn.Create(&parent).GetErrors()
	errs = pr.conn.Create(&user).GetErrors()
	return errs
}

func (pr *GormParentRepository) GetParents() ([]models.Parent, []error) {
	var parents []models.Parent
	errs := pr.conn.Find(&parents).GetErrors()
	return parents, errs
}

func (pr *GormParentRepository) GetParentById(id uint) (models.Parent, []error) {
	parent := models.Parent{}
	errs := pr.conn.Where("id = ?", id).First(&parent).GetErrors()
	return parent, errs
}

func (pr *GormParentRepository) DeleteParent(id uint) []error {
	errs := pr.conn.Unscoped().Where("id = ?", id).Delete(&models.Parent{}).GetErrors()
	errs = pr.conn.Unscoped().Where("id = ?", id).Delete(&models.User{}).GetErrors()
	return errs
}

func (pr *GormParentRepository) UpdateParent(newParent models.Parent) (models.Parent, []error) {
	parent := models.Parent{}
	errs := pr.conn.Model(&parent).Where("id = ?", newParent.Id).Updates(&models.Parent{Email: newParent.Email, Password: newParent.Password}).GetErrors()
	return parent, errs
}

func (pr *GormParentRepository) ViewGrade(parentId uint) (models.Student, []error) {
	student := models.Student{}
	var result []models.Result
	errs := pr.conn.Where("parent_id = ?", parentId).First(&student).GetErrors()
	errs = pr.conn.Where("student_id = ?", student.Id).Find(&result).GetErrors()
	student.Result = result
	return student, errs
}

func (pr *GormParentRepository) GetChild(parentId uint) (models.Student, []error) {
	student := models.Student{}
	errs := pr.conn.Where("parent_id = ?", parentId).First(&student).GetErrors()
	return student, errs
}
