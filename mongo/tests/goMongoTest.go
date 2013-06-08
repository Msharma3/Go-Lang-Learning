package goMongoTest

import (
	"encoding/csv"
	"fmt"
	"github.com/garyburd/go-mongo/mongo"
	"log"
	"os"
)

func main() {
	// create file
	file, err := os.Create("items.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create CSV Writer
	writer := csv.NewWriter(file)

	// Connect to MongoDB
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

	// Auto Detect Headerline
	var m mongo.M
	err = items.Find(nil).One(&m)
	if err != nil && err != mongo.Done {
		log.Fatal(err)
	}

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

		var record []string
		var header []string
		for k, _ := range m {
			s := fmt.Sprint(v)
			record = append(record, s)
		}
		writer.Write(record)
		writer.Flush()
	}
}
