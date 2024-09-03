package main

import (
	"fmt"
	"net/http"
	"os"
)

func main()  {
	res, err := http.Get("https://google.com");

	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}

	fmt.Println(res);

}