package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	var resultUrls []string
	if len(os.Args) != 2 {
		fmt.Println("Provide one url")
		return
	}
	url := os.Args[1]

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if href := scanner.Text(); strings.Contains(href, "href=") {
			for _, val := range strings.Split(href, " ") {
				if strings.Contains(val, "href") {
					val = strings.Split(val, "\"")[1]
					if val != "#" {
						resultUrls = append(resultUrls, val)
					}
				}
			}
		}
	}

	for _, val := range resultUrls {
		fmt.Println(val)
	}
}
