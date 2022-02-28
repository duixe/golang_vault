package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to the Slices section y'all");

	// as close as slices might look like an array, the 
	// only different is that no length is passed in. i.e [2]
	var itemList = []string{"Radio", "Laptop", "Wrist-Watch"};
	fmt.Printf("the type of item list is: %T", itemList);

	fmt.Println("");
	fmt.Println("---------");
	fmt.Println("");

	itemList = append(itemList, "Mug", "Bottle");
	fmt.Println(itemList);  // [Radio Laptop Wrist-Watch Mug Bottle]
	fmt.Println("");

	// getting or printing between ranges using slices
	itemList = append(itemList[1:]);
	fmt.Println(itemList); // [Laptop Wrist-Watch Mug Bottle]

	itemList = append(itemList[1:3]);
	fmt.Println(itemList); // [Laptop Wrist-Watch]

	itemList = append(itemList[:3]);
	fmt.Println(itemList); // [Radio Laptop Wrist-Watch]

	//another format of making a slice through memory allocation
	randomNum := make([]int, 4);
	randomNum[0] = 23;
	randomNum[1] = 90;
	randomNum[2] = 34;
	randomNum[3] = 12;

	fmt.Println("")
	fmt.Println(randomNum);
	fmt.Println("")
	//randomNum[4] = 111; // this will throw an out of bound exception because it uses the current allocated memory
	randomNum = append(randomNum, 45, 56, 887); //append will reallocate another memory to randomNum and add the given values so no out of bound error will occur
	fmt.Println(randomNum);

	fmt.Println("")
	fmt.Println(sort.IntsAreSorted(randomNum));
	//sorting the elements in a slice
	sort.Ints(randomNum);
	
	fmt.Println("sorted Nums: 👉", randomNum);
	fmt.Println(sort.IntsAreSorted(randomNum));

	//removing an element in a slice based on index - using the append method
	var brands = []string {"Dell", "Lenovo", "HP", "Apple", "Samsung"};
	var index int = 2;

	//start from the 0th index up to the defined which is 2, exclude the 3rd index and continue
	brands = append(brands[:index], brands[index+1:]...);
	fmt.Println(brands);

}