# golang url parsing

untuk memparsing string ke bentuk url (`url.URL`), bisa dengan menggunakan package `net/url`.
Dengan memparsing ke bentuk yang sesuai (`url.URL`), ada banyak informasi yang bisa didapatkan,
diantaranya adalah jenis protokol, host, path, dan query.

```go
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
```

function `url.Parse(rawURL string)` digunakan untuk memparsing string ke bentuk url (`*url.URL`),
function ini mengembalikan dua data, yang pertama object yang bertipe `*url.URL` dan error (jika ada).

jika saat memparsing string berbentuk url, jika ada query pada url-nya maka akan di parsing ke bentuk `map[string][]string`, dengan key adalah nama element query, dan value array string yang berisikan value element query.