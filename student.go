package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, s := range students {
		if s.Id == id {
			return s
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{Id: "E001", Name: "ethan",Grade: 2})
	students = append(students, &Student{Id: "E002", Name: "bourne",Grade: 2})
	students = append(students, &Student{Id: "E003", Name: "jason",Grade: 3})
}