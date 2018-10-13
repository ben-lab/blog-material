package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/tv42/httpunix"
)

func handleHTTP(w http.ResponseWriter, req *http.Request) {

	fmt.Printf("Requested : %s\n", req.URL.Path)

	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("docker-socket", "/var/run/docker.sock")

	req.URL.Scheme = "http+unix"
	req.URL.Host = "docker-socket"

	resp, err := u.RoundTrip(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
func main() {

	server := &http.Server{
		Addr:    ":8888",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleHTTP(w, r) }),
	}

	log.Fatal(server.ListenAndServe())
}
