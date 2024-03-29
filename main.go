package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "https://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}

		response, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("status: %s\n", response.Status)
		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
