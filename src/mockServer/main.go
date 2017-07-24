package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asdine/storm"
	"github.com/boltdb/bolt"

	"net/http"

	"goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	db, err := storm.Open("db/my.db", storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8000", mux)
}
