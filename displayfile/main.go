package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	}
	fname := os.Args[1]
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}
	file_data, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(file_data))
}
