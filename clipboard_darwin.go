// +build darwin

package clipboard

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func write(file string) error {
	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf("set the clipboard to (read \"%s\" as TIFF picture)", file))
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(b))
	}
	return nil
}

func read() (io.Reader, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	f.Close()
	defer os.Remove(f.Name())

	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf("write (the clipboard as «class PNGf») to (open for access \"%s\" with write permission)", f.Name()))
	b, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", err, string(b))
	}

	buf := new(bytes.Buffer)
	f, err = os.Open(f.Name())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := io.Copy(buf, f); err != nil {
		return nil, err
	}

	return buf, nil
}
