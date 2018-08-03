package initializer

import (
    "fmt"
    "github.com/Shan4Kites/stringutil/stringutility"
    "net/http"
    "github.com/Shan4Kites/GoShanPersonnal/handlers"
    "os"
    log "github.com/sirupsen/logrus"
)


func initializeLogging() (file *os.File) {
    if _, err := os.Stat("log"); os.IsNotExist(err) {
        os.Mkdir("log", 0755)
    }
    file, err := os.OpenFile("log/production.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    log.SetOutput(file)
    return
}

func Initialize() {
    file := initializeLogging()
    defer file.Close()

    fmt.Printf(stringutility.Reverse("!oG ,olleH"))
    http.ListenAndServe(":8080", handlers.Router())
}
