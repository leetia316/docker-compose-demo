package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("变身超级赛亚人..."))
	})

	// used by docker healthcheck
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`pong`))
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}
