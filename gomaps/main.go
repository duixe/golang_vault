package main

import "fmt"

func main()  {
	fmt.Println("goland maps");
	fmt.Println("")

	brands := make(map[string]string);

	brands["D"] = "Dell";
	brands["LV"] = "Lenovo";
	brands["A"] = "Apple";
	brands["SA"] = "Samsung";

	fmt.Println("list of Element in the brands map: 👉", brands);
	fmt.Println("The key LV has a value of: 👉", brands["LV"]);
	fmt.Println("")

	delete(brands, "A");
	fmt.Println("lis of brands after deleting Apple: 👉", brands);
	fmt.Println("")

	//looping through a map using for loop (foreach in other languages for key value maps)
	for key, value := range brands {
		fmt.Printf("key %v has a value of %v \n", key, value);
	}
}