package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	httpsPort = "3000"
	httpPort  = "8090"
)

func Redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	hostNameParts := strings.Split(req.Host, ":")
	target := "https://" + hostNameParts[0] + ":" + httpsPort + req.URL.Path

	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func main() {
	var port string
	flag.StringVar(&port, "port", httpPort, "http port")
	flag.Parse()

	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 16,
		Handler:        http.HandlerFunc(Redirect)}

	log.Fatal(server.ListenAndServe())
}
