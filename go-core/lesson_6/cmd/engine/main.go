package main

import (
	"bufio"
	"fmt"
	"os"
	"pkg/crawler"
	"pkg/engine"
	"pkg/index"
	"pkg/storage/file"
	"strings"
)

func main() {
	storage, err := file.NewStorage()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Running DB updating in separate process...")
	go collectUpdates(storage)

	eng := engine.New(storage)
	var word string

	fmt.Println("Welcome to the lesson 6")

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

		found, err := eng.Search(strings.ToLower(word))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Results for '%s':\n", word)
		for _, rec := range found {
			fmt.Printf("%s - %s\n", rec.URL, rec.Title)
		}
	}

}

func collectUpdates(storage *file.Storage) {
	crw := crawler.New("https://habr.com", 2)
	webData, err := crw.Scan()
	if err != nil {
		fmt.Println(err)
		return
	}

	ind := index.New(storage)
	err = ind.Fill(&webData)
	if err != nil {
		fmt.Println(err)
		return
	}
}
