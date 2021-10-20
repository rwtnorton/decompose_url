package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing url argument(s)")
		os.Exit(1)
	}
	for _, rawURL := range os.Args[1:] {
		processURL(rawURL)
	}
}

func processURL(rawURL string) {
	var err error
	url, err := url.Parse(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing url: %v\n", err)
		os.Exit(1)
	}
	baseURL := url.Scheme + `://` + url.Host
	fmt.Printf("%s\n", baseURL)
	path := url.Path
	if len(path) != 0 {
		fmt.Printf("  %s\n", path)
	}
	q := url.Query()
	if len(q) > 0 {
		for k, vs := range q {
			var v string
			if len(vs) == 1 {
				v = fmt.Sprintf("%q", vs[0])
			} else {
				v = fmt.Sprintf("%q", vs)
			}
			fmt.Printf("    %q => %s\n", k, v)
		}
	}
}
