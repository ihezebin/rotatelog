//go:build !linux
// +build !linux

package rotatelog

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
