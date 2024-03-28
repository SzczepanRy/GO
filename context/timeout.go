package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"net/http"
)

func main() {

	//sets a timeout conext that gives the req 100 ms
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)

	//cancels timeout
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	imageData, err := io.ReadAll(res.Body)
	fmt.Printf("download image size %d \n", len(imageData))
}
