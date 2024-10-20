package toolkit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTools_DownloadStaticFile(t *testing.T) {
	r := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	var testTool Tools
	testTool.DownloadStaticFile(r, req, "./testdata/ABCD.jpg", "CodeList.jpg")

	res := r.Result()
	defer res.Body.Close()

	if res.Header["Content-Length"][0] != "306633" {
		t.Error("wrong content length of", res.Header["Content-Length"][0])
	}

	if res.Header["Content-Disposition"][0] != "attachment; filename=\"CodeList.jpg\"" {
		t.Error("wrong content disposition")
	}

	_, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
}
