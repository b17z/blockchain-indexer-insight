package routes

import (
	"encoding/json"
    "net/http"
)

type VersionResponse struct {
	Version 	string		`json:"version"`
}

func Version(w http.ResponseWriter, req *http.Request) {
	version := VersionResponse{"0.0.1"}

	js, err := json.Marshal(version)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
