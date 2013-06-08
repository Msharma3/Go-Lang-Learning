package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/garyburd/go-mongo/mongo"
	"log"
	"os"
	"sort"
)

func main() {
	// *host
	host := flag.String("host", "localhost:27017", "The host and port of the mongod instance you wish to connect to")
	dbname := flag.String("db", "testDB", "The output file that will be written to")
	user := flag.String("user", "", "The user you wish to authenticate with")
	pass := flag.String("pass", "", "The pass you wish to authenticate with")
	col := flag.String("col", "testCol", "The collection you wish to output")
	// fieldsFile := flag.String("fieldsFile", "~/", "The a file containing field names")
	// query := flag.String("query", "collection.csv", "The output file that will be written to")
	out := flag.String("out", "collection.csv", "The output file that will be written to")
	flag.Parse()

	// After cmd flag parse
	file, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create Writer
	writer := csv.NewWriter(file)

	// Connect to MongoDB
	conn, err := mongo.Dial(*host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// select DB
	db := mongo.Database{conn, *dbname, mongo.DefaultLastErrorCmd}
	// test for user and pass vars
	err = db.Authenticate(*user, *pass)
	if err != nil {
		log.Fatal(err)
	}

	// select Collection
	collection := db.C(*col)

	// Auto Detect Headerline
	var h mongo.M
	err = collection.Find(nil).One(&h)
	if err != nil {
		log.Fatal(err)
	}

	var headers []string
	for k, _ := range h {
		headers = append(headers, fmt.Sprint(k))
	}
	// Default sort the headers
	// Otherwise accessing the headers will be
	// different each time.
	sort.Strings(headers)
	// write headers to file
	writer.Write(headers)
	writer.Flush()
	// log.Print(headers)

	// Create a cursor using Find query
	cursor, err := collection.Find(nil).Cursor()
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
		for _, header := range headers {
			record = append(record, fmt.Sprint(m[header]))
		}
		writer.Write(record)
		writer.Flush()
	}
}
