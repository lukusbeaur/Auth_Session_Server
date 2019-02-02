package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn
var users = map[string]string{
	"User1": "password1",
	"User2": "password2",
}

func main() {
	initCache()
	//"Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)

	//start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
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

//Welcome func is the controller for the welcome screeen
func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Hello from Welcome Func")
}

//Signin func is the controller for the sign in screen
func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Hello from the Sign in Func")
}
