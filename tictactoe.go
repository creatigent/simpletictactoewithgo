package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	array := [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	var player int

	PrintBoard(array)

	// Example https://thenewstack.io/cli-command-line-programming-with-go/
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}

		// split line args
		args := strings.Split(line, " ")

		// convert string to int
		a1, _ := strconv.Atoi(args[0])
		a2, _ := strconv.Atoi(args[1])

		a1 = a1 - 1
		a2 = a2 - 1

		if len(args) > 2 {
			fmt.Println("Try again...")
			PrintBoard(array)
		} else {
			if player%2 == 0 {
				array[a1][a2] = "X"
				PrintBoard(array)
				player++
			} else {
				array[a1][a2] = "O"
				PrintBoard(array)
				player++
			}
		}
	}
}

// PrintBoard is ...
func PrintBoard(array [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(array[i][j])
			if j != 2 {
				fmt.Print(",")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
