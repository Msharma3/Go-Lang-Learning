package main

import (
	"fmt"
	"github.com/sourcec0de/Go-Lang-Learning/reddit"
	"log"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
