package repository

import (
	"fmt"
	"github.com/jperezviloria/student-go/model"
	"github.com/jperezviloria/student-go/repository/configSql"
)

func GetAllStudents() []model.Student {

	var allStudents []model.Student
	studentRepository := configSql.ConnectSqlServer()
	studentRepository.Raw("SELECT * FROM Students").Scan(&allStudents)
	return allStudents
}

func GetStudentById(idStudent int) model.Student {

	var particularStudent model.Student
	studentRepository := configSql.ConnectSqlServer()
	studentRepository.Raw(fmt.Sprintf("SELECT * FROM Student WHERE id = '%d'", idStudent)).Scan(&particularStudent)
	return particularStudent
}

func SaveStudent(student model.Student) string {

	studentRepository := configSql.ConnectSqlServer()
	studentRepository.Raw(fmt.Sprintf("INSERT INTO Student (name, age) VALUES ('%s','%d')", student.Name, student.Age))
	return "student was saved"
}
