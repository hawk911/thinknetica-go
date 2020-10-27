package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"strings"
)

var words string

func main() {

	url := "https://habr.com"
	fmt.Println("Scanning... ", url)
	const depth = 2
	titles, err := crawler.Scan(url, depth)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Welcome to the lesson 2")
	for {
		fmt.Println("Enter some words that its need to find or leave empty for exit:")
		words := ""
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

		for k, v := range titles {
			if strings.Contains(lower(k), lower(words)) || strings.Contains(lower(v), lower(words)) {
				fmt.Printf("%s - '%s'\n", k, v)
			}
		}
	}

}

func lower(str string) string {
	strLower := strings.ToLower(str)
	return strLower
}
