package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(os.Args[1], nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "go app " + os.Args[1])
}
