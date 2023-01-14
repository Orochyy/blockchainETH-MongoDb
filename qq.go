package main

import (
	"fmt"
	"net/http"
)

func main() {
	qq, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(qq)
}
