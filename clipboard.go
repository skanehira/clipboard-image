package clipboard

import (
	"io"
	"io/ioutil"
	"os"
)

// CopyToClipboard copy image to clipboard
func CopyToClipboard(r io.Reader) error {
	f, err := writeTemp(r)
	if err != nil {
		return err
	}
	defer os.Remove(f)
	return copyToClipboard(f)
}

// ReadFromClipboard read image from clipboard
func ReadFromClipboard() (io.Reader, error) {
	return readFromClipboard()
}

func writeTemp(r io.Reader) (string, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return "", err
	}
	return f.Name(), nil
}
