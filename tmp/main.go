package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error fetching url:", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Saving response body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0666)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Error writing to file:", file)
			}
		}
	}

}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
	}
	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("-", 20))
	}
}
