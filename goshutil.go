package goshutil

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func Copy(src, dst string) (written int64, err error) {
	// Opening source file.
	srcf, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcf.Close()

	// Obtaining destination info.
	fi, err := os.Stat(dst)
	if err != nil {
		return 0, err
	}

	// Opening destination file.
	var dstf *os.File
	if fi.IsDir() {
		fp := filepath.Join(dst, filepath.Base(src))
		dstf, err = os.Create(fp)
	} else {
		dstf, err = os.Open(dst)
	}

	if err != nil {
		return 0, err
	}

	defer dstf.Close()
	return io.Copy(dstf, srcf)
}

// TODO: Make Windows friendly.
// File determines a file type of the input file path.
func File(fpath string) (string, error) {
	cmd := exec.Command("file", fpath)
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	output := string(bytes.TrimSpace(b))
	return output, nil
}
