package toolkit

import (
	"os"
	"testing"
)

func TestTools_CreateDirIfNotExist(t *testing.T) {
	err := CreateDirIfNotExist("../testdata/myDir")
	if err != nil {
		t.Error(err)
	}
	err = CreateDirIfNotExist("../testdata/myDir")
	if err != nil {
		t.Error(err)
	}

	os.Remove("./testdata/myDir")
}
