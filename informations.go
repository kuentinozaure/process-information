package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type ServerInformation struct {
	Os        string
	ProcState string
	Date      string
}

func getServerInfo(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	contents, err := ioutil.ReadFile("/proc/stat")
	dt := time.Now()
	if err != nil {
		return
	}

	m := ServerInformation{os, string(contents), dt.String()}
	js, err := json.Marshal(m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func main() {
	http.HandleFunc("/", getServerInfo)
	http.ListenAndServe(":8080", nil)
}
