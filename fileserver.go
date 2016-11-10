package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

const usageText = `Usage: fileserver [-a host:port] [pattern:]dir ...

The program serves files from a local dist using dir argument as root.
An optional pattern specified the URL base path. For instance,
to serve /tmp dir at /temporary, one can use the following:

    fileserver /temporary:/tmp

The dir argument can be specified multiple times.
If no dir argument is provided, current directory is used.

`

var address *string = flag.String("a", "localhost:8000", "Address (host:port) to listen on")

func main() {
	flag.Usage = func() {
		os.Stderr.WriteString(usageText)
		flag.PrintDefaults()
	}
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = append(roots, ".")
	}
	for _, r := range roots {
		var p string
		if i := strings.IndexRune(r, ':'); i >= 0 {
			p = r[:i]
			r = r[i+1:]
		}
		if !strings.HasPrefix(p, "/") {
			p = "/" + p
		}
		fs := http.FileServer(http.Dir(r))
		if len(p) > 1 {
			fs = http.StripPrefix(p, fs)
		}
		http.Handle(p, fs)
	}

	log.Printf("serving %s on http://%s", roots, *address)
	log.Fatal(http.ListenAndServe(*address, http.HandlerFunc(handle)))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)
	http.DefaultServeMux.ServeHTTP(w, r)
}
