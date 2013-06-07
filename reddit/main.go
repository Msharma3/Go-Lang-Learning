// All Go code belongs to a package.
// Go programs begin with function main inside package main
package main

// Each string is an import path. It tells Go tools where to find
// the package. These package are all from the Go Standard Library.
import (
	"io"
	"log"
	"net/http"
	"os"
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
	// Use io.Copy to copy the HTTP response body to standard output(os.Stdout)
	_, err = io.Copy(os.Stdout, resp.Body)
	// The resp.Body type implements to io.Reader and os.Stdout implements io.Writer

	if err != nil {
		log.Fatal(err)
	}
}
