package logintest

import (
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var tools toolkit.Tools
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&payload); err != nil {
		tools.ErrorJSON(c.Writer, err) // Adjust this if you have a custom error handler in your toolkit
		return
	}

	var resPayload toolkit.JSONResponse
	if payload.Username == "me@here.com" && payload.Password == "verysecret" {
		resPayload.Error = false
		resPayload.Message = "Logged in"
		tools.WriteJSON(c.Writer, http.StatusAccepted, resPayload) // Adjust to use c.Writer for response
		return
	}

	resPayload.Error = true
	resPayload.Message = "invalid credentials"
	tools.WriteJSON(c.Writer, http.StatusAccepted, resPayload) // Adjust to use c.Writer for response
}
