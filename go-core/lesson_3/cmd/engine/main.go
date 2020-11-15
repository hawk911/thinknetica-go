package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/engine"
)

func main() {
	url := "https://habr.com/"
	const depth = 2
	var words string
	crw := crawler.New(url, depth)
	fmt.Println("Scanning... ", url)

	fmt.Println("Welcome to the lesson 3")
	for {
		fmt.Println("Enter some words that its need to find or leave empty for exit:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			words = scanner.Text()
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Something wrong, try again in a few minutes, please, :", err)
			return
		}

		if len(words) == 0 {
			fmt.Println("Thank you to use our solution!")
			break
		}

		found, err := engine.Search(crw, words)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Results for '%s':\n", words)
		for _, a := range found {
			fmt.Println(a)
		}
	}

}
