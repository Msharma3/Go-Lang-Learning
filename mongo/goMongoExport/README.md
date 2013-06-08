# goMongoExport

This is a script writen in go to support exporting a collection from a MongoDB database to a CSV file.

You can build this using the Go build tools, and install it as a cross platform command line tool.

After installation you can simply run this command via command line passing the these flags and it will dynamically discover headers for the CSV and write them to the specified file.

~~~shell
goMongoExport -host $HOST -db $DBNAME -user $USER -pass $PASS -col $COLLECTION -out $OUTPUT_FILE
~~~


I will continue to improve this tool, as the mongo export tool that comes with mongoDB can be a little annoying when exporting
to anything other than JSON. CSVs usually contain BSON objects like _id:ObjectId("..."), and ISODate("...").

This is an attempt to make a cleaner export to a well formatted CSV file.
Also to allow anyone to just add this executable, run it, and have a dump of their DB in CSV format.


Note* Remote DB's may take a while to download due to network connectivity, and latency.
But the result will be just as nice unless interupted.