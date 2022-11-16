package server

import (
	"encoding/json"
	"fmt"
	"godevil/src/hardware/disk"
	"io"
	"net"
	"net/http"
)

var (
	contentJSON string = "application/json"
)

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", contentJSON)
	// (*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "GET, POST")
	(*w).Header().Add("Access-Control-Allow-Headers", "Accept, Contrnt-Type, Content-Length, Authorization, X-CSRF-Token, Accept-Encoding")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Printf("got /devices GET request\n")
		o := disk.Get()
		json.NewEncoder(w).Encode(&o)
	case http.MethodPost:
		fmt.Printf("got /devices POST request\n")
		w.WriteHeader(http.StatusBadRequest)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Middleware to perform operations before working on the actual request
func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setHeaders(&w)
		handler.ServeHTTP(w, r)
	})
}

func getMux() *http.ServeMux {
	deviceHandle := http.HandlerFunc(deviceHandler)

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/devices", middlewareHandler(deviceHandle))
	mux.Handle("/devices/", middlewareHandler(deviceHandle))
	return mux
}

func NewHttp(host string, port uint16) error {
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	fmt.Printf("Webserver listening on %s to accept requests...\n", address)

	mux := getMux()
	err := http.ListenAndServe(address, mux)

	if err != nil {
		panic(err)
	}
	return nil
}
