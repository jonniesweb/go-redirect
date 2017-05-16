// based off of https://gist.github.com/hSATAC/5343225
// modified slightly to accept command line parameters

package main

import (
	"log"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, os.Args[2], 301)
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: redirect <port number> <url>\n" +
			"eg. redirect 8080 https://google.com")
	}

	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":"+os.Args[1], nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
