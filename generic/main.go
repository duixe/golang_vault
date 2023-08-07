package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
);

func main()  {
	fmt.Println("Welcome to the File refresher");

	var content string = "This text will surely be moved into a text file";

	file, err := os.Create("./sample.txt");

	handleError(err)

	length, err := io.WriteString(file, content);

	handleError(err);

	fmt.Println("Total character contained in a file is: ", length);
	defer file.Close();

	readFile("./sample.txt");
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename);

	handleError(err);

	fmt.Println("Data bytes of the parsed file: ", string(databyte));
}

func handleError(err error) {
	if (err != nil) {
		panic(err);
	}
}