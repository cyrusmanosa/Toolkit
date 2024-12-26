package toolkit

import (
	"os"
	"testing"
)

func TestTools_CreateDirIfNotExist(t *testing.T) {
	var testTools Tools
	err := testTools.CreateDirIfNotExist("../testdata/myDir")
	if err != nil {
		t.Error(err)
	}
	err = testTools.CreateDirIfNotExist("../testdata/myDir")
	if err != nil {
		t.Error(err)
	}

	os.Remove("./testdata/myDir")
}
