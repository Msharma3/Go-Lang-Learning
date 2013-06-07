#Go Reddit JSON Reader

Redit JSON Response Example
~~~json
{
	"data":{
		"children":[
			{
				"data":{
					"title":"The Go Homepage",
					"url": "http://golang.org"
				}
			}
		]
	}
}
~~~
Go's json package decodes JSON-encoded data into native Go data structures. To decode the API response, decalre types that reflect the structure of the JSON data.
~~~go
type Item struct {
	Title string
	URL string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}
~~~



Instead of copying the HTTP response body to standard output
~~~go
_, err = io.Copy(os.Stdout, resp.Body)
~~~~
we use the json package to decode the response into our Response data structure.
~~~go
r := new(Response)
err = json.NewDecoder(resp.Body).Decode(r)
~~~
Initilize a new Response value, store a pointer to it in the new variable r. Create a new json.Decoder object and decode the response body into r.

As the decoder parses the JSON data it looks for corresponding fields of the same names in the Response struct. The "data" field of the top-level JSON object is decoded into the Response struct's Data field, and the JSON array "children" is decoded into the Children slice, and so on.


Iterate over the Children slice, assigning the slice value to the child on each iteration. The Prinln call prints the item's Title followed by a newline.

~~~go
for _, child := range r.Data.Children {
	fmt.Println(child.Data.Title)
}
~~~

NOTE* _, means you are ignoring a value returned from a function while assigning variables.