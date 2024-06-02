package server

import (
	applerefurbished "apple-refurbished/src/lib"
	"fmt"
	"log"
	"net/http"
)

const ServerPort = 8080

func Handler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	params := r.URL.Query()
	url := params.Get("url")
	bucket := params.Get("bucket")
	filename := params.Get("filename")

	if bucket == "" {
		http.Error(w, "missing bucket", http.StatusBadRequest)
		return
	}
	var rUrl string
	if url != "" {
		rUrl = url
	} else {
		rUrl = applerefurbished.DefaultUrl
	}

	var filepath string
	var err error
	if filename != "" {
		filepath, err = applerefurbished.DumpWithFilename(rUrl, bucket, filename)
	} else {
		filepath, err = applerefurbished.Dump(rUrl, bucket)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to dump: %v", err), http.StatusInternalServerError)
	}

	// Respond with a greeting
	response := fmt.Sprintf("Data available under: 'gs://%s/%s'", bucket, filepath)
	w.Write([]byte(response))
}

func Run() {
	http.HandleFunc("/", Handler)

	log.Printf("Starting server on :%d", ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
