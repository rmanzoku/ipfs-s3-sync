package ipfss3

import (
	"bytes"
	"io"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var ipfsUrl = os.Getenv("IPFS")

func LoadIpfs(path string) ([]byte, error) {
	sh := shell.NewShell(ipfsUrl)
	r, err := sh.Cat(path)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}

func StoreIpfs(body []byte) (hash string, err error) {
	sh := shell.NewShell(ipfsUrl)
	r := bytes.NewReader(body)
	return sh.Add(r)
}
