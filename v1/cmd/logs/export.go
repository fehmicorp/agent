package logs

import (
	"io"
	"os"
)

func Export(destination string) error {

	src, err := os.Open(LogFile)
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	return err
}
