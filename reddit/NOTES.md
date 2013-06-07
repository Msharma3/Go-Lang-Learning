Go Reddit JSON Reader

<!-- Redit JSON Response Example -->
{
	"data":{
		"children":[
			{
				"data":{
					"title":"The Go Homepage",
					"url": "http://golang.org",
					...
				}
			},
			...
		]
	}
}

Go's json package decodes JSON-encoded data into native Go data structures. To decode the API response, decalre types that reflect the structure of the JSON data.