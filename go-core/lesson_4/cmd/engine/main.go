package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/invertindex"
	"strings"
)

func main() {
	url := "https://habr.com/"
	const depth = 2
	var word string
	crw := crawler.New(url, depth)
	fmt.Println("Scanning... ", url)
	i := invertindex.New(crw)
	err := i.Fill()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Welcome to the lesson 4")
	for {
		fmt.Println("Enter a word that its need to find or leave empty for exit:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			word = scanner.Text()
			break
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Something wrong, try again in a few minutes, please, :", err)
			return
		}

		if len(word) == 0 {
			fmt.Println("Thank you to use our solution!")
			break
		}

		found := i.Search(strings.ToLower(word))

		fmt.Printf("Results for '%s':\n", word)
		for _, a := range found {
			fmt.Println(a)
		}
	}

}
