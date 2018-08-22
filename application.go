package main

import (
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"github.com/rs/xid"
	"github.com/Shan4Kites/GoShanPersonnal/test"
	"github.com/onrik/logrus/filename"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	justChecking(r.Context())
}

func justChecking(ctx context.Context) {
	log := log.WithFields(log.Fields{"requestId": ctx.Value("requestId")})
	log.Info( "Url is called now")
	test.TestChange(ctx)
}

func RequestMetaDataInitializer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestId := xid.New()
		ctx := context.WithValue(req.Context(), "requestId", requestId)
		req = req.WithContext(ctx)
		log := log.WithFields(log.Fields{"requestId": ctx.Value("requestId")})
		log.Info("Before")
		h.ServeHTTP(w, req) // call original
		log.Println("After")
	})
}

func TestWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Inside TestWrapper Before")
		h.ServeHTTP(w, r) // call original
		log.Println("Inside TestWrapper After")
	})
}

func initializeLogging() (file *os.File) {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", 0755)
	}
	file, err := os.OpenFile("log/production.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(file)
	log.AddHook(filename.NewHook())
	return
}

func main() {
	file := initializeLogging()
	defer file.Close()
	log.Println("Service just started")
	r := mux.NewRouter()
	r.HandleFunc("/", Handler)

	http.Handle("/", RequestMetaDataInitializer(r))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
