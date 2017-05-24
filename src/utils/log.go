package utils

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type RequestLogger struct {
	Handle http.Handler
	Logger *log.Logger
}

func (rl RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rl.Handle.ServeHTTP(w, r)
	log.Printf("[%s] %s %s in %v", strings.Replace(os.Args[0], "./", "", -1), r.Method, r.URL.Path, time.Since(start))
}
