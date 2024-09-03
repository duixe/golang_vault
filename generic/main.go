package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("Let's talk file operations in golang");

	/*
	 create a string e.g "This text will be tranfered to a file"
	 create a file
	 write the string to the created file
	 read from the file
	
	*/

	content := "😃 Make sure you smile even in the storm";

	File, err := os.Create("sample.txt");

	errorHandler(err);

	length, contentErr := io.WriteString(File, content);

	errorHandler(contentErr);

	fmt.Println("The length of this file is: ", length);

	defer File.Close();
	readFileContent("sample.txt");

}

func errorHandler(err error) {
	if err != nil {
		panic(err);
	}
}


func readFileContent(filename string) {
	dataByte, err := os.ReadFile(filename);

	errorHandler(err);

	fmt.Println("The content of the file is: \n", string(dataByte))
}
	



