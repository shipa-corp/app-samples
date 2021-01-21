package main

import (
	"fmt"
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
	// port is hardcoded and shipa takes it from shipa.yaml
	port := "9999"
	c := conf{
		port:        port,
		username:    u.Username,
		dir:         dir,
	}
	http.HandleFunc("/", configHandler(c))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

type conf struct {
	port        string
	username    string
	dir         string
}

func configHandler(c conf) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "port: %s\n", c.port)
		fmt.Fprintf(w, "user: %s\n", c.username)
		fmt.Fprintf(w, "working directory: %s\n", c.dir)
	}
}
