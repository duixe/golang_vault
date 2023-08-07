package main

import "fmt"

func main()  {
	fmt.Println("Welcome toi the func section");

	spitFactsTwo();
	fmt.Println("--")
	spitFacts();

	//NB: Golang does not support higher level function - wriiting a function in another function
	// hence this 👇 will through an error
	// func thisFuncIsntRIght()  {
	// 	bla bla bla
	// }
	result := addNUms(3, 6);

	fmt.Println("Sum of two int equals: ", result);

	fmt.Println("")

	getResult := addUnknownAmountOfInts(3,6,9,1,2,9);

	fmt.Println("Adding numerous results gives: ", getResult);

	fmt.Println("");

	resultInt, resultString := returnIntAndString(2, 3, 5,6, 8);
	fmt.Println("ResultInt is: ", resultInt);
	fmt.Println("resultString is: ", resultString);

}

func spitFacts()  {
	fmt.Println("Faith without works is vague");
}

func spitFactsTwo()  {
	fmt.Println("One's destiny can never be destroyed, it can only be delayed");
}

func addNUms(numOne int, numTwo int) int  {
	return numOne + numTwo;
}

// Add unknown amount of integers
// NB: values as an attribute becomes a slice after adding the three dots
func addUnknownAmountOfInts(values ...int) int {
	total := 0;

	for _, value := range values {
		total += value;
	}

	return total;
}

//functions that returns more than one type
func returnIntAndString(values ...int) (int, string)  {
	total := 0;
	var resultString string = " Greetings";

	for _, value := range values {
		total += value;
	}

	return total, resultString;
}