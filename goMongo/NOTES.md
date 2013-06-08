#MongoDB

It seems go has two drivers that are not
officially supported by 10Gen.

- mgo
- Go-mongo

`Go-mongo` appears to have a simpler interface for mongodb, where as `mgo` is more complex, but for good reason. `mgo` comes packed with additional features like failover, and cluster discovery. However, this may not be for you if what you are trying to achieve is a simple mongo interface connecting to a single mongod instance.

In this case I would suggest using `Go-mongo`.

I will return to the subject of mongoDB and Go once I mastered more of the basics of go, and after I have created some csv files using the standard go library.