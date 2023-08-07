package main

import "fmt"

//NB: The := operator cannot be used outside a function e.g main()

//NB: A global variable must start with capital letter
const GlobalVar string = "globavar";

func main()  {
	//decalring a variable with different types

	//string
	var nickname string = "duixe";
	fmt.Printf("I'm Emmanuel but my nickname is %s and it's a type of %T \n ", nickname, nickname);

	//Boolean
	var isAdult bool = true;
	fmt.Printf("variable is of type %T \n", isAdult);

	//integers
	// to avoid any misunderstanding one can specify int instead of uint8,16 etc.
	var smallValue uint8 = 252;
	fmt.Printf("variable %d is of type %T \n", smallValue, smallValue);

	//floats
	var smallDecimals float32 = 255.789076544323456;
	var bigDecimals float64 = 255.789076544323456;

	fmt.Printf("variable decimals %v is of type: %T \n", smallDecimals, smallDecimals);
	fmt.Printf("variable decimals %v is of type: %T \n", bigDecimals, bigDecimals);

	//decalaring variables and aliases
	var declaredVar int;
	fmt.Printf("variable %v as been decalred but not instantiated \n", declaredVar); //unassigned int vars returns 0

	//decalring by excluding return types
	var stringVar = "this has been assigned a string value, hence can not be changed to another type";
	fmt.Println(stringVar + " " +GlobalVar);
	// stringVar = 3 x - this is not accepted since the initial value assigned to the variable is of type string

	// decalring by excluding var completing and accepting 
	numberofTuts := 870;
	fmt.Println(numberofTuts);
}