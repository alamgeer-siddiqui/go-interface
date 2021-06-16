package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("ERROR:=", err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	// bs:=make([]byte,99999)
	// resp.Body.Read(bs)
	// fmt.Println(string (bs))

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bite", len(bs))
	return len(bs), nil
}
