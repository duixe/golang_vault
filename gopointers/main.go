package main

import "fmt"

func main()  {
	fmt.Println("welcome to the pointers section");

	//creating a pointer add "*" before the data type of the variable
	//var pointer *string;

	//fmt.Println("The value of this pointer is: ", pointer); //The value of this pointer is:  <nil>

	// create variable that serves as a pointer to another variable which is already created
	myNum := 33;

	var pointer = &myNum;

	fmt.Println("Address of pointer is: ", pointer);  //Address of pointer is00003XDf
	fmt.Println("Value of pointer is: ", *pointer);  //Value of pointer is: 33

	*pointer = *pointer - 2;
	fmt.Println("value of the actual variable - myNum is: ", myNum);  // value of the actual variable myNum is: 31


}