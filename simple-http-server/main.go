package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")

	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	return ip
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("有人来了")

	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	os.Setenv("VERSION", "0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	clientIP := getCurrentIP(r)
	httpCode := http.StatusOK
	log.Printf("ip && httpCode:", clientIP, httpCode)

	w.Write([]byte("hello, world"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerIndex)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("error:", err.Error())
	}
}
