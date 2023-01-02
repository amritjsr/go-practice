package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World!!!! Greetings from Amritjsr")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your value : ")
	input, _ := reader.ReadString('\n')
	aInt, err := strconv.ParseInt(strings.TrimSpace(input), 10,64 )
	
	if err != nil {
		fmt.Printf("Some Problem occured : ")
		fmt.Println(err)
	} else {
		fmt.Printf("All is well : ")
		fmt.Println(aInt)
	}
}