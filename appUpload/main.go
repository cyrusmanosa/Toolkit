package main

import (
	"fmt"
	"log"
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit"
)

func main() {
	mux := routes()

	log.Println("Starting server...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/upload", uploadFiles)
	mux.HandleFunc("/upload-one", uploadOneFile)

	return mux
}

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"image/png", "image/jpeg", "image/gif"},
	}

	files, err := t.UploadFiles(r, "./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	out := ""
	for _, item := range files {
		out += fmt.Sprintf("Uploaded %s to the uploads folder to %s\n", item.OriginalFileName, item.NewFileName)
	}

	_, _ = w.Write([]byte(out))
}

func uploadOneFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"image/png", "image/jpeg", "image/gif"},
	}

	files, err := t.UploadOneFile(r, "./uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Uploaded 1 file, %s, to the uploads folder", files.OriginalFileName)))
}
