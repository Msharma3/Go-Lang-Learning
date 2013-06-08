package main

import (
	"encoding/csv"
	// "encoding/json"
	"fmt"
	"os"
)

func main() {
	// Create a writeable file
	file, err := os.Create("items.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create CSV Writer
	writer := csv.NewWriter(file)

	// Create Header Row
	headers := []string{"item_name", "brand_id"}
	err = writer.Write(headers)
	writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

	// Example
	// for _, obj := range d {
	// var record []string
	// record = append(record, strconv.FormatInt(obj.RecordID, 10))
	// record = append(record, obj.DOJ)
	// record = append(record, obj.EmpID)
	// w.Write(record)
	// }
	// w.Flush()

}
