package main

import (
	"fmt"
	"net/url"
)

const urlString string = "https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=json&dummy=nothing";

func main()  {
	fmt.Println("Handling URLs in Golang");

	//parsing a url with golang url package
	res, _ := url.Parse(urlString);

	//Extracting the url scheme
	// fmt.Println(res.Scheme);

	//Extracting the url Host
	// fmt.Println(res.Host);

	//Extracting the url Path
	// fmt.Println(res.Path);

	//Extracting the url scheme
	// fmt.Println(res.RawQuery);

	/*
		NB: when the url has more query params and one is trying to access
		go lang url package already has a method that returns all the params
		in a map, hence you can perfrom a an iteration on them
	*/

	fmt.Printf("Type of query params is: %T \n", res.Query());

	for i, v := range res.Query() {
		fmt.Printf("query %v is %v \n", i, v);
	}


	//constructing a url;
	var urlPath = &url.URL{
		Scheme: "https",
		Host: "akomaning.com",
		Path: "/home",
		RawPath: "user=duixe",
		RawQuery: "seearch=portfolio",
	}

	newlyFormedUrl := urlPath.String();
	fmt.Println(newlyFormedUrl);
}