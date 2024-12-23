package scanner

import (
	"os"

	"github.com/pkg/errors"
	"github.com/randyttruong/tla2typst/pkg/util"
)

type Loader struct {
	buf    string
	bufLen int
}

var (
	loader = &Loader{}
)

func GetLoader() *Loader {
	return loader
}

func SetBuffer(arr string) {
	loader.buf = arr
	loader.bufLen = len(loader.buf)
}

func LoadDocument(filepath string) error {

	err := util.CheckFilePermissionsAndOwnership(filepath)

	if err != nil {
		errors.Wrapf(err, "Failed to load document, instead got: %v", err)
	}

	bytes, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.Wrapf(err, "File at %v does not exist. Existing", filepath)
		}

		return errors.Wrapf(err, "Failed to read file, got %v", err)
	}

	if bytes == nil {
		err = errors.New("Bytes is empty")
		return err
	}

	SetBuffer(string(bytes))

	return nil
}
