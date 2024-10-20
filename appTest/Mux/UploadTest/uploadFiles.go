package uploadtest

import (
	"fmt"
	"net/http"

	toolkit "github.com/cyrusmanosa/Toolkit/v2"
)

func UploadFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	t := toolkit.Tools{
		MaxFileSize:      1024 * 1024 * 1024,
		AllowedFileTypes: []string{"application/pdf"},
	}

	files, err := t.UploadFiles(r, "/Users/cyrusman/Desktop/ProgrammingLearning/Udemy/Toolkit/appTest/Mux/UploadTest/uploads")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	out := ""
	for _, item := range files {
		out += fmt.Sprintf("Uploaded %s to the uploads folder to %s\n", item.OriginalFileName, item.NewFileName)
	}

	_, _ = w.Write([]byte(out))
}
