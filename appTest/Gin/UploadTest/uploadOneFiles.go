package uploadtest

import (
	"fmt"
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
	"github.com/gin-gonic/gin"
)

func UploadOneFiles(c *gin.Context) {
	t := toolkit.Tools{
		MaxFileSize: 1024 * 1024 * 1024,
		AllowedFileTypes: []string{
			"application/pdf",
			"image/jpeg", // For JPEG and JPG
			"image/png",  // For PNG
		},
	}

	files, err := t.UploadOneFile(c.Request, "/Users/cyrusman/Desktop/ProgrammingLearning/Udemy/Toolkit/appTest/Gin/UploadTest/uploads")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Uploaded 1 file, %s, to the uploads folder", files.OriginalFileName))
}
