package main

import "fmt"

func main()  {
	fmt.Println("Welcome to the struct section of my goland vault") // usually refer to as an alternative for classes in other programming languages
	//NB: NO inheritance in golang

	employee := Employee{"Duixe", 27, 34443, "duixe@dev.co", true};
	fmt.Println(employee);
	fmt.Printf("Employee details are: 👉 %+v\n", employee);

	fmt.Println("");
	
	fmt.Printf("Employee %v email is %v \n", employee.Name, employee.Email);

	
}

type Employee struct {
	Name string;
	Age int;
	staffNumber int64;
	Email string;
	Active bool;
}