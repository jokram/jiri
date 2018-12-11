// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textutil

import (
    "golang.org/x/crypto/ssh/terminal"
    "os"
)

// TerminalSize returns the dimensions of the terminal, if it's available from
// the OS, otherwise returns an error.
func TerminalSize() (row, col int, _ error) {
	// Try getting the terminal size from stdout, stderr and stdin respectively
	// using terminal.GetSize function, which should work on most platforms.
	
	fdOut := int(os.Stdout.Fd())
	fdIn := int(os.Stdin.Fd())
	fdErr := int(os.Stderr.Fd())
	
	if row, col, err := terminal.GetSize(fdOut); err == nil {
		return row, col, err
	}
	if row, col, err := terminal.GetSize(fdErr); err == nil {
		return row, col, err
	}
	return terminal.GetSize(fdIn)
}
