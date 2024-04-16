package classes

type Employee struct {
	Name          string
	Age           int
	StreetAddress string
	PostalCode    int
	City          string
}

func (emp Employee) CalculateSalary() int {
	return emp.Age * 100
}

// func CalculateSalary(emp Employee) int {
// 	return emp.Age * 100
// }

func New(name string, age int) Employee {
	return Employee{Name: name, Age: age}
}

// class Employee{
// 	Name string
// 	City string

// 	float CalculateSalary(){}
// }

// e = new Employee();
// e.Name = "Stefan";
// s:= e.CalculateSalary();
