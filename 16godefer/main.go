package main

import "fmt"

func main()  {
	defer fmt.Println("world");
	defer fmt.Println("one");
	defer fmt.Println("two");
	fmt.Println("Hello");
	deferFunc();
	//the outcome this program will be as follows

	//Hello 4 3 2 1 0 two one world
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v \n", i);
	}
}