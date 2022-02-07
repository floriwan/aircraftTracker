package handler

import (
	"aircraftTracker/acdb"
	"aircraftTracker/observer"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type stats struct {
	Aircrafts int
	Observer  int
	Enroute   int
}

func HandleAircraftReg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reg := mux.Vars(r)["reg"]
	switch r.Method {
	case "PUT":
		err := observer.Add(reg)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf(`{"message": "%v"}`, err)))
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		err := observer.Remove(reg)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf(`{"message": "%v"}`, err)))
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "not implemented"}`))
	}
}

func GetStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "not implemented"}`))
	}

	b, _ := json.Marshal(stats{Aircrafts: acdb.GetSize(), Observer: observer.GetSize(), Enroute: 0})

	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func SearchAircraft(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "not implemented"}`))
	}

	owner := r.URL.Query()["owner"]

	// convert string slice int string and search for this owner
	l := acdb.SearchOwner(strings.Join(owner, ""))
	b, err := json.Marshal(l)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "error %v"}`, err)))
	}

	log.Printf("search for owner %v, return %v aircraft(s)\n", owner, len(l))

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetObservationList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	l := observer.GetObservationList()
	var buffer bytes.Buffer
	for k := range l {
		buffer.WriteString("'")
		buffer.WriteString(l[k].Reg)
		buffer.WriteString("' ")
	}
	w.Write(buffer.Bytes())

}

func GetAircraftData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "not implemented"}`))
	}

	reg := mux.Vars(r)["reg"]
	log.Printf("request aircraft registration '%v'\n", reg)
	acInfo, err := acdb.GetAcInfo(reg)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(acInfo)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}
