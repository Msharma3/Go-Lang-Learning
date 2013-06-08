package main

import (
	"github.com/garyburd/go-mongo/mongo"
	"log"
	// "os"
	// "time"
	// "fmt"
)

func main() {
	conn, err := mongo.Dial("dharma.mongohq.com:10053")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Wrap the connection with a logger so that we can view traffic
	// to and from the mongoDB server
	// conn = mongo.NewLoggingConn(conn, log.New(os.Stdout, "", 0), "")
	// clear log prefix for more readable output
	log.SetFlags(0)

	// Create db object
	db := mongo.Database{conn, "nix_stagingv3", mongo.DefaultLastErrorCmd}
	err = db.Authenticate("sync", "vbnvbn45")
	if err != nil {
		log.Fatal(err)
	}

	// Create collections object
	items := db.C("items")

	// Create a cursor using Find query
	cursor, err := items.Find(nil).Cursor()
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close()
	// Iterate over all items in a collection
	for cursor.HasNext() {
		var m mongo.M
		err := cursor.Next(&m)
		if err != nil {
			log.Fatal(err)
		}
		// Iterate over all keys in returned
		// doc and log the value of the map key
		for k, _ := range m {
			log.Print(k, ": ", m[k])
		}
	}
	// expectedFieldValues(cursor, "item_name")

}
