package school

type Student struct {
	ID 					int
	FirstName, LastName string
	Age 				int
	Class				int
	Subject				[]string
	Grade				map[string]string
}

var Subject  = map[int][]string{
	JSS1: { "English", "Int. Science", "Intro-Tech"},
	JSS2: { "English", "Int. Science", "Intro-Tech"},
	SSS1: { "Chemistry", "Biology", "Physics"},
	SSS2: { "Chemistry", "Biology", "Physics"},
	SSS3: { "Chemistry", "Biology", "Physics"},
}


var Students = map[int]Student{
	1: {
		ID: 1,
		FirstName: "Clin",
		LastName:  "Ade",
		Age:       15,
		Class:     SSS3,
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

var Alumni = 10


