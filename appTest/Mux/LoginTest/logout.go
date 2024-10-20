package logintest

import (
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "Logged out",
	}

	tools.WriteJSON(w, http.StatusAccepted, payload)
}
