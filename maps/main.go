package main

import "fmt"

func main() {
	// countries := map[string]string{
	// 	"Africa": "Ghana",
	// 	"Europe": "France",
	// }

	// fmt.Println(countries);

	//Another method of declaring a map
	// countries := make(map[string]string);

	// countries["Africa"] = "Ghana";
	// countries["Asia"] = "India";

	// delete(countries, "Asia");
	// fmt.Println("countries: ", countries);
	countries := map[string]string{
		"africa": "Ghana",
		"asia": "japan",
		"europe": "England",
		"north america": "canada",
	};

	iterate(countries);
}

func iterate(variable map[string]string)  {
	for i, v := range variable {
		fmt.Printf("continent %v has %v \n", i, v);
	}
}