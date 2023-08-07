package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	hello := "welcome to input output section";
	fmt.Println(hello);

	fmt.Println("");

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the Rating for our course: ");

	// comma ok || comma err - also known as try catch in other programming language
	// i.e when one is expecting an error - _, error := "this format will surely produce an err"
	input, _ := reader.ReadString('\n');
	fmt.Println("Thanks for your rating us", input);
	fmt.Printf("The type of rating is: %T \n", input);
}