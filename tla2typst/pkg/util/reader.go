package util

import (
	"fmt"
	"os"
	"os/user"
	"syscall"

	"github.com/pkg/errors"
)

func CheckFilePermissionsAndOwnership(filepath string) error {
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
