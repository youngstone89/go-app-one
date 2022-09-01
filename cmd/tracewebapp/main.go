package main

import (
	"net/http"
	_ "net/http/pprof"
)

// curl localhost:8181/debug/pprof/trace?seconds=10 > trace.out
// wrk -c 100 -t 10 -d 60s http://localhost:8181/hello

func main() {
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	http.ListenAndServe("localhost:8181", http.DefaultServeMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}
