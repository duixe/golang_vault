package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	var hello string = "welcome to the typecasting section";

	fmt.Println(hello);
	fmt.Println("PLease kindly rate my course from 1 to 7");

	reader := bufio.NewReader(os.Stdin);

	input, _ := reader.ReadString('\n');

	fmt.Println("Thanks for rating us with ", input);

	//try and get the input (in string and covert it);
	ratedNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64);

	if err != nil {
		fmt.Println("oops: ", err)
	}else {
		fmt.Println("Thanks, oo and we've added one to your rating: ", ratedNum+1);
	}
}