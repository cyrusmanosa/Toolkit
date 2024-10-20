package uploadtest

import (
	"fmt"
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
	"github.com/gin-gonic/gin"
)

func UploadFiles(c *gin.Context) {
	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"application/pdf"},
	}

	files, err := t.UploadFiles(c.Request, "/Users/cyrusman/Desktop/ProgrammingLearning/Udemy/Toolkit/appTest/Gin/UploadTest/uploads")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	out := ""
	for _, item := range files {
		out += fmt.Sprintf("Uploaded %s to the uploads folder to %s\n", item.OriginalFileName, item.NewFileName)
	}

	c.String(http.StatusOK, out)
}
