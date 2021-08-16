package school

type Student struct {
	ID 					int
	FirstName, LastName string
	Age 				int
	Class				int
	Subject				[]string
	Grade				map[string]string
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


