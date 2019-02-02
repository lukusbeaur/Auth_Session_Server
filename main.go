package main

import (
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn

func main() {
	initCache()
	//"Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)

	//start the server
	log.Fatal(http.ListenAndServe("8080", nil))
}

func initCache() {
	//initialize the redis connection
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}

	//assign the connection to the package level
	cache = conn
}

func Welcome(w http.ResponseWriter, r *http.Request) {

}
func Signin(w http.ResponseWriter, r *http.Request) {

}
