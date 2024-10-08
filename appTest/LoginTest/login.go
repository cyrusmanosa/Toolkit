package logintest

import (
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
