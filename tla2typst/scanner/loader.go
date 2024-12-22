package scanner

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

type Loader struct {
	buf []string
}

var (
	loader = &Loader{}
)

func GetLoader() *Loader {
	return loader
}

func SetBuffer(arr []string) {
	loader.buf = arr
}

func checkFilePermissionsAndOwnership(filepath string) error {
	fileInfo, err := os.Stat(filepath)

	if err != nil {
		return errors.Wrapf(err, "Failed to stat file: %v", err)
	}

	fileMode := fileInfo.Mode().Perm()

	if fileMode != 0o600 {
		return errors.Wrapf(err, "File does not have 600 permissions, got: %v", fileMode)
	}

	fileOwner, ok := fileInfo.Sys().(*syscall.Stat_t)

	if !ok {
		return errors.Wrapf(err, "Failed to get system info")
	}

	fileOwnerUid := fmt.Sprintf("%d", fileOwner.Uid)

	currentUser, err := user.Current()

	if err != nil {
		return errors.Wrapf(err, "Failed to get current user: %v", err)
	}

	if fileOwnerUid != currentUser.Uid {
		return errors.Wrapf(err, "File is not owned by current user, instead: %v", fileOwnerUid)
	}

	return nil
}

func LoadDocument(filepath string) error {

	err := checkFilePermissionsAndOwnership(filepath)

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

	// TODO: Write the actual parser

	return nil
}
