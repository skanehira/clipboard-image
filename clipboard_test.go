package clipboard

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyToClipboard(t *testing.T) {
	r, err := os.Open(filepath.Join("testdata", "out.png"))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	if err := Write(r); err != nil {
		t.Fatal(err)
	}
}

func TestReadFromClipboard(t *testing.T) {
	r, err := Read()
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if len(b) == 0 {
		t.Fatal("clipboard data length is 0")
	}
}
