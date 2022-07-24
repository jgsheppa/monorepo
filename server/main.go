package main

import (
	"encoding/json"
	"fmt"
	"github.com/jgsheppa/monorepo/net"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port  = "8080"
	feUrl = "http://localhost:3000"
)

type IP struct {
	Address string `json:"address"`
}

type Domain struct {
	Address []net.IP `json:"address"`
}

func IpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", feUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	vars := mux.Vars(r)
	ip := vars["address"]
	w.WriteHeader(http.StatusOK)

	address := IP{
		Address: ip,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)

}

func DomainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", feUrl)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	vars := mux.Vars(r)
	ip := vars["ip"]

	addr, err := network.LookupDomain(ip)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	address := Domain{
		Address: addr,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ip/{address}", IpHandler)
	r.HandleFunc("/domain/{ip}", DomainHandler)

	// HandlerFunc converts notFound to the correct type
	fmt.Println("Starting the development server on port" + port)
	http.ListenAndServe(": "+port, r)
}
