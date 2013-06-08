package mongo

import (
	"encoding/csv"
	"fmt"
	"github.com/garyburd/go-mongo/mongo"
	"log"
	"os"
)

// A function that Dials and Authenticates
// into a provided data base
// accespts host,db,user,pass,log=Bool
// returns db,err
func ConnAndAuth(h, d, u, p string, l bool) mongo.Database {
	conn, err := mongo.Dial(h)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if l {
		conn = mongo.NewLoggingConn(conn, log.New(os.Stdout, "", 0), "")
		// clear log prefix for more readable output
		log.SetFlags(0)
	}

	db := mongo.Database{conn, d, mongo.DefaultLastErrorCmd}
	err = db.Authenticate(u, p)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Dynacmiaclly create headers for a csv
// file by getting setting the field names
// into a slice
// accepts a collection
// returns collection, err
func CreateHeadersDyn(c mongo.Collection) []string {
	var m mongo.M
	err := c.Find(nil).One(&m)
	if err != nil && err != mongo.Done {
		log.Fatal(err)
	}

	var header []string
	for k, _ := range m {
		header = append(header, fmt.Sprint(k))
	}
	return header
}

// creates a file, and a writer object
// returns writer,err
func CreateWriter(fn string) (w csv.Writer) {
	file, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w = csv.NewWriter(file)
	return w
}

// writes a record
// to the file, and flushes the write
func WriteRow(r []string, w csv.Writer) {
	w.Write(r)
	w.Flush()
}
