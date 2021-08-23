package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
	log.Println(resp.Body)

	// --------------------------------
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// log.Println(string(bs))
	// --------------------------------

	// --------------------------------
	// io.Copy(os.Stdout, resp.Body)
	// --------------------------------
	lw := &logWriter{}
	io.Copy(lw, resp.Body)
}

func (l *logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
