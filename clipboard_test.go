package clipboard

import (
	"bytes"
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

	if err := CopyToClipboard(r); err != nil {
		t.Fatal(err)
	}
}

func TestReadFromClipboard(t *testing.T) {
	r, err := ReadFromClipboard()
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	o, err := ioutil.ReadFile(filepath.Join("testdata", "out.png"))
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(b, o) != 0 {
		t.Fatalf("clipboard data length: %d, test data length: %d", len(b), len(o))
	}
}
