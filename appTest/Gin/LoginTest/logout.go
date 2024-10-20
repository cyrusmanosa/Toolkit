package logintest

import (
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "Logged out",
	}

	// Use tools.WriteJSON with Gin's context writer
	tools.WriteJSON(c.Writer, http.StatusAccepted, payload)
}
