// +build darwin

package clipboard

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func copyToClipboard(file string) error {
	cmd := exec.Command("osascript", "-e",
		fmt.Sprintf("set the clipboard to (read \"%s\" as TIFF picture)", file))
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(b))
	}
	return nil
}

func readFromClipboard() (io.Reader, error) {
	return nil, errors.New("mac os doesn't support")
}
