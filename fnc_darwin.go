// Copyright 2019 TempleX (temple3x@gmail.com).
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package fnc

import (
	"os"
	"syscall"
)

// Disable all File access time(atime) updates,
// darwin doesn't have it.
const O_NOATIME = 0

func syncRange(f *os.File, off int64, n int64, flags int) (err error) {

	// on darwin even call Sync, the drive may not write dirty page to the media.
	// It's not a big deal here, because darwin is only for test environment.
	// Here sync all data.
	_, _, errno := syscall.Syscall(syscall.SYS_FCNTL, f.Fd(), uintptr(syscall.F_FULLFSYNC), uintptr(0))
	if errno == 0 {
		return nil
	}
	return errno
}

func fadvise(f *os.File, offset, size int64, advice int) (err error) {
	return
}

func preAllocate(f *os.File, size int64) error {
	return f.Truncate(size)
}
