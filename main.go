package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", healthzHandleFunc)
	http.HandleFunc("/", defaultHandleFunc)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzHandleFunc(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	log.Println(GetIP(request), 200)
}

func defaultHandleFunc(writer http.ResponseWriter, request *http.Request) {
	_, err := ioutil.ReadAll(request.Body)
	var statusCode int
	if err != nil {
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}
	for k, v := range request.Header {
		writer.Header()[k] = v
	}
	writer.Header()["VERSION"] = []string{os.Getenv("VERSION")}
	writer.WriteHeader(statusCode)
	log.Println(GetIP(request), statusCode)
}

func GetIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "-"
	}

	if net.ParseIP(ip) != nil {
		return ip
	}

	return "-"
}
