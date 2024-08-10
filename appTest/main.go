package main

import (
	"log"
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
)

func main() {
	mux := routes()

	log.Println("Starting application on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	mux.HandleFunc("/api/login", login)
	mux.HandleFunc("/api/logout", logout)

	return mux
}

func login(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := tools.ReadJSON(w, r, &payload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	var resPayload toolkit.JSONResponse
	if payload.Username == "me@here.com" && payload.Password == "verysecret" {
		resPayload.Error = false
		resPayload.Message = "Logged in"
		tools.WriteJSON(w, http.StatusAccepted, resPayload)
		return
	}

	resPayload.Error = true
	resPayload.Message = "invalid credentials"
	tools.WriteJSON(w, http.StatusAccepted, resPayload)
}

func logout(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "Logged out",
	}

	tools.WriteJSON(w, http.StatusAccepted, payload)
}
