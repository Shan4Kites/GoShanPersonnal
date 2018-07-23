package initializer

import (
    "fmt"
    "github.com/Shan4Kites/stringutil/stringutility"
    "net/http"
    "github.com/Shan4Kites/GoShanPersonnal/handlers"
)

func Initialize() {
    fmt.Printf(stringutility.Reverse("!oG ,olleH"))
    http.ListenAndServe(":8080", handlers.Router())
}
