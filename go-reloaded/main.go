package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Check numbers of arguments")
		return
	}
	// open file
	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// do something with a word
		(scanner.Text())
		fmt.Print(" ")
	}
	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
