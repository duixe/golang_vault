package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to file section");

	content := "😀😀 .. Well kindly put me inside a txt file 😎😎";

	file, contentErr := os.Create("testFile.txt");

	checkErrorExist(contentErr);

	length, err := io.WriteString(file, content);

	checkErrorExist(err);

	fmt.Println("the length of this file is: ", length);
	// file.Close();
	// recommended way of closing a file is by defering it
	defer file.Close();
	readFile("testFile.txt")
}

func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename);

	checkErrorExist(err);

	fmt.Println("The content of the file is displayed below 👇 \n", string(databyte));
}

func checkErrorExist(err error)  {
	if err != nil {
		panic(err);
	}
}