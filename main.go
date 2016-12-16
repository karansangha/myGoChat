package main

import "flag"
import "fmt"

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		fmt.Println("is host")
	} else {
		fmt.Println("is not host")
	}
}
