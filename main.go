package main

import "fmt"

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Has           []byte
}

func main() {
	fmt.Println("Hello world!")
}
