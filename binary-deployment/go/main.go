package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", HelloServer)
    port := os.Getenv("PORT")
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "binary deployment")
}
