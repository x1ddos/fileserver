package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port    *int    = flag.Int("p", 8000, "Port to listen at")
	address *string = flag.String("a", "localhost", "Address to listen on")
	root    *string = flag.String("d", ".", "Root directory")
)

func main() {
	flag.Parse()

	httproot := http.Dir(*root)
	*address = fmt.Sprintf("%s:%d", *address, *port)
	fmt.Printf("Serving %s on http://%s\n", httproot, *address)

	http.Handle("/", http.FileServer(httproot))
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
