package school

import "github.com/Ad3bay0c/SchoolManagement/students"

var Subject  = map[int][]string{
	JSS1: { "English", "Int. Science", "Intro-Tech"},
	JSS2: { "English", "Int. Science", "Intro-Tech"},
	SSS1: { "Chemistry", "Biology", "Physics"},
	SSS2: { "Chemistry", "Biology", "Physics"},
	SSS3: { "Chemistry", "Biology", "Physics"},
}


var Students = map[int]students.Student{
	1: {
		ID: 1,
		FirstName: "Clin",
		LastName:  "Ade",
		Age:       15,
		Class:     SSS1,
	},
	2: {
		ID:        2,
		FirstName: "Clint",
		LastName:  "Ade",
		Age:       15,
		Class:     SSS2,
	},
	3: {
		ID:        3,
		FirstName: "Clinto",
		LastName:  "Ade",
		Age:       15,
		Class:     SSS3,
	},
}


