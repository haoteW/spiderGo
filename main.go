package main

import "net/http"

func main() {
	resp, err := http.Get(
		"https://movie.douban.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
