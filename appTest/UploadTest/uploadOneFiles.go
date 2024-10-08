package uploadtest

import (
	"fmt"
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
)

func UploadOneFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"application/pdf"},
	}

	files, err := t.UploadOneFile(r, "/Users/cyrusman/Desktop/ProgrammingLearning/Udemy/Toolkit/appTest/UploadTest/uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Uploaded 1 file, %s, to the uploads folder", files.OriginalFileName)))
}
