package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://www.fanialfi.space:80/hello?name=fani alfirdaus&age=17"

	url, err := url.Parse(urlString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s\n\n", urlString)

	fmt.Printf("protocol\t: %s\n", url.Scheme)
	fmt.Printf("host\t\t: %s\n", url.Host)
	fmt.Printf("path\t\t: %s\n", url.Path)
	fmt.Printf("query\t\t: %v\n", url.Query())

	query := url.Query()
	fmt.Printf("\nnama\t: %s\n", query["name"][0])
	fmt.Printf("umur\t: %s\n", query["age"][0])
}
