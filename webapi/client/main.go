package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// @see https://deeeet.com/writing/2016/11/01/go-api-client/

func main() {
	url, _ := url.Parse("https://api.github.com/search/repositories?q=aaa")
	method := "GET"
	body := strings.NewReader("")

	req, _ := http.NewRequest(method, url.String(), body)

	// req.SetBasicAuth("authuser", "authpass")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "testcode-golangClient/0.1 (0.1)")

	client := &http.Client{}

	res, _ := client.Do(req)

	out := new(bytes.Buffer)
	res.Write(out)

	fmt.Println(out)
}
