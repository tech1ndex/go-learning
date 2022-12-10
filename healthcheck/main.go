package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}

	_, err := client.Head("https://hc.tech1ndex.ca/ping/d3603d5e-ba64-4e26-b866-65e1a07d02a6")
	if err != nil {
		fmt.Printf("%s", err)
	}
}
