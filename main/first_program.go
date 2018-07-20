package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/Shan4Kites/stringutil/stringutility"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    fmt.Printf(stringutility.Reverse("!oG ,olleH"))
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
