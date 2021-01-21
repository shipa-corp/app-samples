package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

var port string

func main() {
    http.HandleFunc("/", HelloServer)
    port = os.Getenv("PORT")
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    before, err := ioutil.ReadFile("/tmp/before-hook")
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
    }
    after, err := ioutil.ReadFile("/tmp/after-hook")
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
    }
    fmt.Fprintf(w, "%s%s", before, after)
}
