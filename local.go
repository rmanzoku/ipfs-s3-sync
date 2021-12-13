package ipfss3

import (
	"errors"
	"io"
	"os"
)

func LoadLocal(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func StoreLocal(path string, body []byte) (err error) {

	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, _ = os.Create(path)
		} else {
			return err
		}
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	return nil
}
