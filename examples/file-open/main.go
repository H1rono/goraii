package main

import (
	"fmt"
	"io"

	"github.com/H1rono/goraii"
)

func main() {
	for file := range goraii.OpenFile("README.md") {
		readme, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(readme))
	}
}
