package models

type Employee struct {
	EmployeeId int `db:"employeeid"`
	FirstName string `db:"firstname"`
	Last_name string `db:"lastname"`
	Position string `db:"position"`
	Department string `db:"department"`
	HireDate string `db:"hiredate"`
	Salary int `db:"salary"`
	Email string `db:"email"`
	PhoneNumber int `db:"phonenumber"`
	Address string `db:"address"`
}
