package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/user"
)

func main() {
    u, err := user.Current()
    if err != nil {
        panic(err)
    }
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    // port is hardcoded and exposed in Dockerfile
    port := "5001"
    c := conf{
        port:     port,
        username: u.Username,
        dir:      dir,
    }
    http.HandleFunc("/", configHandler(c))
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

type conf struct {
    port     string
    username string
    dir      string
}

func configHandler(c conf) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        before, err := ioutil.ReadFile("/home/shipa/before-hook")
        if err != nil {
        	w.WriteHeader(http.StatusBadRequest)
        	return
        }
        after, err := ioutil.ReadFile("/home/shipa/after-hook")
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "port: %s\n", c.port)
        fmt.Fprintf(w, "user: %s\n", c.username)
        fmt.Fprintf(w, "working directory: %s\n", c.dir)
        fmt.Fprintf(w, "%s%s", before, after)
    }
}
