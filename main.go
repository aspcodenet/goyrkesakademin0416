package main

import (
	"fmt"
	"main/classes"

	"github.com/cheynewallace/tabby"
)

// I GO -> alltid copy by value

// func addEmployee(employees []employee) {
// 	employees[0].Name = "3443443"
// 	employees = append(employees, employee{Name: "Josefine"})
// }

// SLICE -> struct length + pekare till en array

func main() {

	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.AddLine("Stefan", "Developer", "Whatever")
	t.AddLine("John Cena", "Wrestler", "US")
	t.Print()

	// constructorer, mandartory namn, age
	emp := classes.New("Stefan", 52)
	emp2 := classes.New("Oliver", 15)
	emp2.PostalCode = 122
	emp.City = "addsdsa"
	//salary := classes.CalculateSalary(emp)
	salary := emp2.CalculateSalary()
	fmt.Println(salary)

	fmt.Println("sdadsadsa")

	// int i;
	//int *i = new int(); // free, malloc

	// players := []employee{
	// 	employee{Name: "Stefan"},
	// 	employee{Name: "Oliver"},
	// }
	//addEmployee(players)
	// fmt.Println(players)

}
