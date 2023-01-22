# go-bookstore

## Inspiration
The project is practically copied from the [The freeCodeCamp.org youtube channel](https://youtu.be/jFfo23yIWac?t=4034). The video contains 11 golang projects.

## What is the project
The project is an api run on ```localhost:3000``` with a few endpoints that a local mysql server. The project you would have a mysql server and enter replace the details of that server in the /pkg/config/app.go file on line 13(Further instruction in the app.go source file). Ther are 5 requests you can make to the server, namely :
* GET ALL @ GET ```/book/``` 
* GET BY ID @ GET ```/book/{bookID}```
* CREATE @ POST ```/book/```
* UPDATE @ PUT ```/book/{bookID}```
* DELETE @ DELETE ```/book/{bookID}```
