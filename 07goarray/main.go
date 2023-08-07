package main

import "fmt"

func main()  {
	fmt.Println("It's time for some array lectures");

	var itemList [4]string;

	itemList[0] = "bottle";
	itemList[1] = "keys";
	itemList[3] = "table";

	//NB: THE third item was skipped
	//NB: Note that the length of the list will always print the exact number specified
	// when decalring the array

	fmt.Println("Item List contains: ", itemList);
	fmt.Println("the length of itemList is: ", len(itemList));  // the length of itemList is: 4

	var brandList = [6]string {"Dell", "HP", "Lenovo", "Apple", "Samsung"};
	fmt.Println("the length of itemList is: ", len(brandList));  // the length of itemList is: 6

}