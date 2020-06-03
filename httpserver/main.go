package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type AnalysisPhase string

type Measurement struct {
	// Phase is the status of this single measurement
	Phase AnalysisPhase `json:"phase"`
	Message string `json:"message,omitempty"`
	Value string `json:"value,omitempty"`
}

type result struct {
	Measurement Measurement `json:"measurement,omitempty"`
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/metrics/", askHandler)
	_ = http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(200)
	log.Println(r.RequestURI)
	log.Println(r.Header)
}

func askHandler(w http.ResponseWriter, r *http.Request) {
	m := new(result)
	if r.URL.Path != "/metrics/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	rand.Seed(time.Now().UnixNano())
	min := 89
	max := 100
	m.Measurement.Value = strconv.Itoa(rand.Intn(max-min +1) + min)
	log.Println("measurement:", m.Measurement.Value)
	b, _ := json.Marshal(m)
	_, _ = w.Write(b)
	log.Println(r.RequestURI)
	log.Println(r.Header)

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		log.Println("Invalid request!!.")
	}
}
