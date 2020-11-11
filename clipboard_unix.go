// +build freebsd linux netbsd openbsd solaris dragonfly

package clipboard

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func copyToClipboard(file string) error {
	b, err := exec.Command("file", "-b", "--mime-type", file).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(b))
	}

	// b has new line
	cmd := exec.Command("xclip", "-selection", "clipboard", "-t", string(b[:len(b)-1]))
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(in, f); err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

func readFromClipboard() (io.Reader, error) {
	cmd := exec.Command("xclip", "-selection", "clipboard", "-t", "image/png", "-o")
	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		return nil, err
	}

	if err := r.Close(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return buf, nil
}
