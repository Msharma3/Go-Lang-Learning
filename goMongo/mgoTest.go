package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Item struct {
	Upc string
}

func main() {
	session, err := mgo.Dial("mongodb://sync:vbnvbn45@dharma.mongohq.com:10053/nix_stagingv3")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("nix_stagingv3").C("items")
	result := Item{}
	err = c.Find(bson.M{"upc": "00000000000001"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Upc:", result.Upc)
}
