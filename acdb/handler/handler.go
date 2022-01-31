package handler

import (
	"aircraftTracker/acdb"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAircraftData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
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

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "not implemented"}`))
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
