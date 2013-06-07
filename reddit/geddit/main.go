package main

import (
	"flag"
	"fmt"
	"github.com/sourcec0de/Go-Lang-Learning/reddit"
	"log"
)

func main() {
	rname := flag.String("r", "golang", "The name of a specified reddit")
	flag.Parse()
	items, err := reddit.Get(*rname)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
