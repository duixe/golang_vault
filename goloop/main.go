package main

import "fmt"

func main()  {
	fmt.Println("Loop section ----");

	names := []string {"john", "doe", "lizo", "jake", "sunny"};

	fmt.Println(names);
	fmt.Println("");

	//for loops - first format
	// for n := 0; n < len(names); n++ {
	// 	fmt.Println(names[n]);
	// }

	// for loops - second format (forr)
	// for n := range names {
	// 	fmt.Println(names[n]);
	// }

	//foreach equivalent for goland
	// for index, name := range names {
	// 	fmt.Printf("index of %v is %v \n", name, index);
	// } 

	//immitating while loop in other languages - NB for keyword is still used
	randomValue := 1;

	for randomValue < 10 {
		// if randomValue == 4 {
		// 	break;
		// }

		if randomValue == 2 {
			goto duixe;
		}

		if randomValue == 5 {
			randomValue++;
			continue;
		}

		fmt.Println("random value: ", randomValue);
		randomValue++;
	}

	// using labels inside loops
	duixe:
		fmt.Println("Print this label out");
}