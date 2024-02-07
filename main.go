package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, World!")

	w.Write([]byte("Shubham Mishra"))

	w.WriteHeader(200)

	//adding headder

	w.Header().Add("Shubham", "Mishra,Google")

	// it will returns an empty header as we have not set any values to it.  map[]
	fmt.Println(w.Header())

	//to know which method request coming from client
	fmt.Println(r.Method)

	//it will return URL of client request not complete URL for this ''http://localhost:8080/helloss'' it will return /helloss
	fmt.Println(r.URL)

	//The protocol version for incoming server requests.

	fmt.Println(r.Proto)

}

func main() {

	http.HandleFunc("/", helloHandler)

	//log.Fatal() is equivalent to Print followed by a call to os.Exit(1).
	// 	Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
	// For portability, the status code should be in the range [0, 125].

	log.Fatal(http.ListenAndServe(":8080", nil))

}
