// All Go code belongs to a package.
// Go programs begin with function main inside package main
package main

// Each string is an import path. It tells Go tools where to find
// the package. These package are all from the Go Standard Library.
import (
	"fmt"
	// "io"
	"encoding/json"
	"log"
	"net/http"
	// "os"
)

// This is a function declaration. The main function takes
// no arguments and has no return values.
func main() {
	// Call the Get function from the http package, passing the URL
	// of the Reddit API as its only argument.
	resp, err := http.Get("http://www.reddit.com/r/golang.json")
	// declare two variables (resp, err) and give them the return values
	// of the function all. (Yes, Go functions can return multiple values.)
	// The Get function returns *http.Response and an error values.

	// Compare err against nil, the zero-value for the built-in error type.
	// The err variable will be nil if the request was successful.
	if err != nil {
		// If not, call the log.Fatal function to print the error
		// message and exit the program.
		log.Fatal(err)
	}

	// // Test that the HTTP server returned a "200 OK"
	if resp.StatusCode != http.StatusOK {
		// If not, bail, printing the HTTP status message ("500 Internal Server Error", for example)
		log.Fatal(resp.Status)
	}

	// Declare Item
	// type Item struct {
	// 	Title string
	// 	URL   string
	// }
	// Declare Response
	type Response struct {
		Data struct {
			Children []struct {
				Data struct {
					Title string
					URL   string
				}
			}
		}
	}

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)

	for _, child := range r.Data.Children {
		fmt.Println(child.Data.Title)
	}

	if err != nil {
		log.Fatal(err)
	}
}
