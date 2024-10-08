package main

import (
	loginTest "appTest/LoginTest"
	uploadTest "appTest/UploadTest"
	"log"
	"net/http"
)

func main() {
	// mux := LoginRoutes()
	mux := UploadRoutes()

	log.Println("Starting application on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func LoginRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./LoginTest"))))

	mux.HandleFunc("/api/login", loginTest.Login)
	mux.HandleFunc("/api/logout", loginTest.Logout)

	return mux
}

func UploadRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./UploadTest"))))

	mux.HandleFunc("/upload", uploadTest.UploadFiles)
	mux.HandleFunc("/upload-one", uploadTest.UploadOneFiles)

	return mux
}
