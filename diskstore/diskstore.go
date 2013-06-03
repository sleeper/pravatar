package diskstore

import (
	// "io"
	"log"
	"os"
	"path"
	"strings"
)

type DiskStore struct {
	Dir string
}

func NewStore(dir string) *DiskStore {
	return &DiskStore{Dir: dir}
}

func (s *DiskStore) Get(hash string) (*os.File, error) {

	var name = hash
	if !strings.HasSuffix(name, ".jpg") {
		name = name + ".jpg"
	}
	var file, err = os.Open(path.Join(s.Dir, name))
	if err != nil {
		log.Fatal(err)
	}
	return file, err
}
