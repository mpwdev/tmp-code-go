package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	result := strings.Contains("I love Go", "ov") // true
	p(result)

	result = strings.ContainsAny("success", "xy") // false - x or y not in success
	p(result)

	result = strings.ContainsRune("I love Go", 'v')
	p(result)

	n := strings.Count("cheese", "e") // 3
	p(n)

	p(strings.ToLower("Hello, World!"))
	p(strings.ToUpper("Hello, World!"))

	p("go" == "go")
	p("Go" == "go")
	p(strings.ToLower("GO") == strings.ToLower("go")) // not good

	p(strings.EqualFold("GO", "go"))
}
