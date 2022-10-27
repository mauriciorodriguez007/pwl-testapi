package main

import (
    "fmt"
    "net/http"
	"time"
)

func main() {
	addr := ":7777"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%sPWL-testapi new request from %v body: %v",time.Now(),r.RemoteAddr,r.Body)
        fmt.Fprintf(w, "%s HELLO WORLD, you've requested: %s\n",time.Now(),r.URL.Path)
    })
	fmt.Printf("\n%sPWL-testapi listening on %v",time.Now(),addr)
    err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("\n%s listenandserve failed error: %s",time.Now(),err)	
	}	
}