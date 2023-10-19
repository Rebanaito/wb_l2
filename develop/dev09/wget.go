package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	argv := os.Args[1:]
	if len(argv) == 0 {
		fmt.Fprintln(os.Stderr, "wget: missing URL")
	}
	for i, target := range argv {
		if !strings.Contains(target, "https://") {
			target = "https://" + target
		}
		req, err := http.NewRequest(http.MethodGet, target, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		filename := "index.html"
		if i != 0 {
			filename += strconv.Itoa(i)
		}
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
