package main

import (
	"fmt"
	"github.com/rs/xid"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"time"
	_ "net/http/pprof"
)

type WelcomeHandler struct {

}
func (WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Our Site!")
}

type LearningHandler struct {

}
func (LearningHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To Learning Page!")
}

func AddRequestIdMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestId := xid.New()
		ctx := context.WithValue(req.Context(), "requestId", requestId)
		req = req.WithContext(ctx)
		h.ServeHTTP(w, req) // call original
	})
}

func LogTimeMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] Time taken is %v\n", r.Context().Value("requestId"), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

func initializeLogging() (file *os.File) {
	file, err := os.OpenFile("/tmp/demo.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(file)
	return
}

func test(i int) {
	fmt.Println("i is :", i)
	time.Sleep(time.Minute)
	fmt.Println("completed")
}

func main() {
	file := initializeLogging()
	defer file.Close()

	var i int
	for i=0; i<10; i++ {
		go test(i)
	}
	http.Handle("/", AddRequestIdMiddleware(LogTimeMiddleware(WelcomeHandler{})))
	http.Handle("/learn", AddRequestIdMiddleware(LogTimeMiddleware(LearningHandler{})))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
