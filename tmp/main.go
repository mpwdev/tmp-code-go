package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
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
	wg.Done()

}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		// fmt.Println(strings.Repeat("-", 20))
	}
	wg.Wait()
}
