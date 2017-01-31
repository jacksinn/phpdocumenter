package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
)

func main() {
	file, err := os.Open("test.php")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var functions []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		matched, _ := regexp.MatchString("function *", scanner.Text())
		if matched {
			fmt.Println("Found function: ", scanner.Text())
			functions = append(functions, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}