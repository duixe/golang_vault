package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://cat-fact.herokuapp.com/facts/";

func main() {
	fmt.Println("Welcome to the http request section");

	res, err := http.Get(url);

	handleError(err);

	fmt.Printf("Reponse type is: %T \n", res);

	/*	
		Always keep in mind that it is the engineer's/developer responsibility
		to always close the http connection when using the net/http module
	*/
	defer res.Body.Close();

	var databytes []byte = handleRequest(res);

	content := string(databytes);

	fmt.Println(content);

}

func handleError(err error) {
	if err != nil {
		panic(err);
	}
}

func handleRequest(res *http.Response) []byte {
	databytes, err := ioutil.ReadAll(res.Body);

	handleError(err);

	return databytes
}