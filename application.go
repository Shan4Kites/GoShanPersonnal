package main

import (
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	justChecking()
}

func justChecking() {
	log.Println("Url is called now")
}

func initializeLogging() (file *os.File) {
	file, err := os.OpenFile("production.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(file)
	return
}

func main() {
	file := initializeLogging()
	defer file.Close()
	log.Println("This is a test log entry")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}