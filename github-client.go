package main

import (
	"fmt"
	"io"
	"net/http"
)

func Connect() {
	resp, err := http.Get("https://api.github.com/users/crl-n")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return
	}

	fmt.Println(string(body))
}
