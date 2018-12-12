package osutil

import (
	"os"
	"os/exec"
	"syscall"
)

func Rename(src, dst string) error {
	if err := os.Rename(src, dst); err != nil {
		// Check if the rename operation failed
		// because the source and destination are
		// located on different mount points.
		linkErr, ok := err.(*os.LinkError)
		if !ok {
			return err
		}
		errno, ok := linkErr.Err.(syscall.Errno)
		if !ok || errno != syscall.EXDEV {
			return err
		}
		// TODO: Check if the code below is reached on Windows
		// if Temp directory is on another drive than the working directory for JIRI
		// Fall back to a non-atomic rename.
		cmd := exec.Command("mv", src, dst)
		return cmd.Run()
	}
	return nil
}
