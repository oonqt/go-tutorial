package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, resp.ContentLength)

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	lw := logWriter{}

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}