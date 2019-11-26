package usecase

import "github.com/nattigy/parentschoolcommunicationsystem/models"

func MakeNewPost(task models.Task) bool {
	return true
}

func EditPost(task models.Task) bool {
	return false
}

func RemoveTask(task models.Task) bool {
	return false
}

func UploadResource(subject models.Subject) {

}

func TeacherUpdateProfile() {

}

func ReportGrade(grade models.Grade) {

}

func ViewClasses() []models.Student {
	return nil
}