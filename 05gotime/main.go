package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Time section of my golang vault");

	currentTime := time.Now();
	fmt.Println(currentTime);

	//printing with regards to go static format
	fmt.Println(currentTime.Format("01-02-2006"));

	// specifying a day
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"));

	//other satatic formats
	fmt.Println(currentTime.Format("Mon Jan 2 15:04:05 MST 2006"));

	//CREATING A CUSTOM TIME AND DATE
	customDateTime := time.Date(1844, time.August, 23, 12, 45, 55, 12, time.UTC);

	fmt.Println(customDateTime.Format("01-02-2006 Monday"));
}