package main

import (
	"encoding/json"
	"net/http"

	"github.com/skeletonkey/lib-core-go/logger"
)

func main() {
	logger := logger.Get()
	logger.Trace().Msg("Start Service")
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// Simulate API response
		data := []byte(`{"message": "Hello from Go!"}`)
		w.Write(data)
	})

	http.Handle("/", calc)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		logger.Debug().Err(err).Msg("ListenAndServe ended poorly")
	}
	logger.Trace().Msg("Service Shutdown")
}

type numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type numsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
	Sub float64 `json:"sub"`
	Div float64 `json:"div"`
}

func process(numsdata numbers) numsResponseData {
	var numsres numsResponseData
	numsres.Add = numsdata.Num1 + numsdata.Num2
	numsres.Mul = numsdata.Num1 * numsdata.Num2
	numsres.Sub = numsdata.Num1 - numsdata.Num2
	numsres.Div = numsdata.Num1 / numsdata.Num2

	return numsres
}

func calc(w http.ResponseWriter, request *http.Request) {
	logger := logger.Get()
	decoder := json.NewDecoder(request.Body)

	var numsData numbers
	var numsResData numsResponseData

	decoder.Decode(&numsData)

	numsResData = process(numsData)
	logger.Printf("numsResData: %v\n", numsResData)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
		logger.Error().Err(err).Msg("something broke during encoding")
	}
}