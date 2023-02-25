package uno

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Go_http_dev() {
	res, err := http.Get("https://go.dev/doc/")
	if err != nil {
		log.Fatal(err)
	}
	results, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", results)
}
