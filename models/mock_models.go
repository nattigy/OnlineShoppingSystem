package models

var ClassRoom1 = ClassRoom{
	Id:         200,
	GradeLevel: 12,
	HomeRoom:   40,
}

var Comment1 = Comment{
	FirstName: "Nati",
	Data:      "My comment",
	TaskId:    2,
	StudentId: 20,
}

var Message1 = Message{
	Id:             10,
	TeacherId:      40,
	ParentId:       30,
	MessageContent: "Hello teacher",
	From:           "parent",
}

var Parent1 = Parent{
	Id:         30,
	FirstName:  "endale",
	MiddleName: "kebede",
	Email:      "endu@gmai.com",
	Password:   "1234",
}

var Result1 = Result{
	Id:         1,
	SubjectId:  100,
	StudentId:  20,
	Assessment: 10,
	Test:       20,
	Final:      30,
	Total:      100,
}

var Student1 = Student{
	Id:          20,
	FirstName:   "yonatahn",
	MiddleName:  "endale",
	Email:       "yoni@gmail.com",
	Password:    "1234",
	SectionId:   300,
	ClassRoomId: 200,
	ParentId:    30,
}

var Subject1 = Subject{
	Id:          100,
	ClassRoomId: 200,
	SubjectName: "Math",
}

var Task1 = Task{
	Id:               1,
	Title:            "fake title",
	Description:      "fake description",
	ShortDescription: "fake ShortDescription",
	SubjectId:        100,
	ClassRoomId:      200,
}

var Teacher1 = Teacher{
	Id:          40,
	FirstName:   "aman",
	MiddleName:  "belete",
	Email:       "aman@gmail.com",
	Password:    "1234",
	ProfilePic:  "",
	SubjectId:   100,
	ClassRoomId: 200,
}

var User1 = User{
	Id:       20,
	Email:    "nati@gmail.com",
	Password: "1234",
	Role:     "student",
	LoggedIn: true,
}
