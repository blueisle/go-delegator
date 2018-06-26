// Package ipfs is a IPFS interface
//
// https://ipfs.io/docs/
package ipfs

import (
	"bytes"
	"io/ioutil"
	"sync"

	"github.com/ipfs/go-ipfs-api"
)

// Ipfs is a IPFS API manager
type Ipfs struct {
	s *shell.Shell
}

// For singleton
var instance *Ipfs
var once sync.Once

// GetInstance returns an instance of Ipfs
func GetInstance(url string) *Ipfs {
	ns := shell.NewShell(url)
	return &Ipfs{
		s: ns,
	}
}

// Cat returns data from IPFS with path(file hash)
func (ipfs *Ipfs) Cat(path string) (ret string) {
	rc, err := ipfs.s.Cat(path)
	if err != nil {
		return
	}
	if b, err := ioutil.ReadAll(rc); err == nil {
		ret = string(b)
		rc.Close()
	}
	return
}

// Add returns path(file hash) after adding data to IPFS
func (ipfs *Ipfs) Add(data string) (string, error) {
	ndata := bytes.NewBufferString(data)
	return ipfs.s.Add(ndata)
}
