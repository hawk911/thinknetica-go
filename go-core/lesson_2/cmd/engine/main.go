package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pkg/crawler"
	"strings"
)

func main() {

	url := reader("Enter link like https://habr.com: \n")
	fmt.Println("Scanning...")
	const depth = 2
	titles, err := crawler.Scan(url, depth)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Welcome to the lesson 2")
	for {
		fmt.Println("Enter some words that its need to find or leave empty for exit:")
		reader := bufio.NewReader(os.Stdin)
		words, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		words = strings.TrimSuffix(words, "\r\n")

		if len(words) == 0 {
			fmt.Println("Thank you to use our solution!")
			break
		}

		for k, v := range titles {
			if strings.Contains(k, words) || strings.Contains(v, words) {
				fmt.Printf("%s - '%s'\n", k, v)
			}
		}
	}

}

func reader(text string) string {
	textStdin := ""
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	//Использую Scan  - попробовал поработать
	for scanner.Scan() {
		textStdin = scanner.Text()
		textStdin = strings.TrimSpace(textStdin)
		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Something wrong, try again in a few minutes, please, :", err)
	}
	if len(textStdin) < 3 {
		log.Fatalln("Param is empty or there are less than 3 symbols!")
	}

	return textStdin
}
