package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Welcome to case and switch in golang");

	rand.Seed(time.Now().UnixNano());
	diceNum := rand.Intn(6) + 1;

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1, kindly open");
	case 2:
		fmt.Println("kindly move 2 times");
	case 3:
		fmt.Println("kindly move 3 times");
		fallthrough
	case 4:
		fmt.Println("kindly move 4 times");
	case 5:
		fmt.Println("kindly move 5 times");
		fallthrough
	case 6:
		fmt.Println("kindly move 6 times and roll again");
	default:
		fmt.Println("wait is that a dice number 😳");
	}
}