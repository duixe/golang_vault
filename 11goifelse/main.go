package main

import "fmt"

func main()  {
	fmt.Println("Control flow section - if else");

	var age int = 18;
	var result string;

	if age < 18 {
		result = "Hey you not eligible to vote 😥";
	}else if age >= 18 {
		result = "Hey you can vote 😀";
	}else {
		result = "Well this means you are'nt born yet - 😎";
	}

	fmt.Println(result);
	fmt.Println("");

	if (12%2 == 0) {
		fmt.Println("12 is an even number");
	}else {
		fmt.Println("12 is an odd number");
	}

	//another syntax where you can initialize a variable and go ahead and apply an if else statement
	if num := 3; num >= 12 {
		fmt.Println("Hey the num I just created is equal or greater than 12");
	}else {
		fmt.Println("oops 3 isn't greater than 12");
	}
}