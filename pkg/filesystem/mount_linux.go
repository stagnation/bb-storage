//go:build linux
// +build linux

package filesystem

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/buildbarn/bb-storage/pkg/filesystem/path"
	"github.com/buildbarn/bb-storage/pkg/util"

	"golang.org/x/sys/unix"
)

const (
	FSCONFIG_SET_FLAG        = 0
	FSCONFIG_SET_STRING      = 1
	FSCONFIG_SET_BINARY      = 2
	FSCONFIG_SET_PATH        = 3
	FSCONFIG_SET_PATH_EMPTY  = 4
	FSCONFIG_SET_FD          = 5
	FSCONFIG_CMD_CREATE      = 6
	FSCONFIG_CMD_RECONFIGURE = 7
)

// Rudimentary function wrapper for the `fsconfig` syscall.
//
// This is implemented as a stop gap solution until real support is merged into the `unix` library.
// See this patchset: https://go-review.googlesource.com/c/sys/+/399995/ .
// This only implements the two commands needed for basic `mountat` functionality.
// And will just exit if any other command is called.
//
// TODO(nils): construct proper `syscall.E*` errors.
//
//	like `unix.errnoErr`, but the function is not exported.
func fsconfig(fsfd int, cmd int, key string, value string, flags int) (err error) {
	switch cmd {
	case FSCONFIG_SET_STRING:
		if len(key) == 0 || len(value) == 0 {
			err = errors.New("`key` and `value` must be provided")
			return
		}
	case FSCONFIG_CMD_CREATE:
		if len(key) != 0 || len(value) != 0 {
			err = errors.New("`key` and `value` must be empty")
			return
		}
	default:
		err = errors.New("not implemented: " + fmt.Sprintf("%d", cmd))
		return
	}

	var _p0 *byte
	var _p1 *byte

	_p0, err = unix.BytePtrFromString(key)
	if err != nil {
		return
	}
	if key == "" {
		_p0 = nil
	}

	_p1, err = unix.BytePtrFromString(value)
	if err != nil {
		return
	}
	if value == "" {
		_p1 = nil
	}

	r0, _, e1 := unix.Syscall6(
		unix.SYS_FSCONFIG,
		uintptr(fsfd),
		uintptr(cmd),
		uintptr(unsafe.Pointer(_p0)),
		uintptr(unsafe.Pointer(_p1)),
		uintptr(flags),
		0,
	)
	ret := int(r0)
	if e1 != 0 {
		err = e1
		return
	}
	if ret < 0 {
		err = errors.New("negative return code, not converted to an error in `Syscall`: " + fmt.Sprintf("%d", ret))
		return
	}

	return
}

// Mounts the `source` filesystem on `mountname` inside a directory
// given as `dfd` file descriptor,
// using the `fstype` filesystem type.
// This returns a file descriptor to the mount object,
// that can be used to move it again.
// Remember to close it before unmounting,
// or unmount will fail with EBUSY.
//
// TODO: Options cannot be sent to the syscalls.
func mountat(dfd int, fstype, source, mountname string) (int, error) {
	fd, err := unix.Fsopen(fstype, unix.FSOPEN_CLOEXEC)
	if err != nil {
		return -1, util.StatusWrapf(err, "Fsopen '%s'", fstype)
	}

	err = fsconfig(fd, FSCONFIG_SET_STRING, "source", source, 0)
	if err != nil {
		return -1, util.StatusWrapf(err, "Fsconfig source '%s'", source)
	}

	err = fsconfig(fd, FSCONFIG_CMD_CREATE, "", "", 0)
	if err != nil {
		return -1, util.StatusWrap(err, "Fsconfig create")
	}

	mfd, err := unix.Fsmount(fd, unix.FSMOUNT_CLOEXEC, unix.MS_NOEXEC)
	if err != nil {
		return -1, util.StatusWrap(err, "Fsmount")
	}
	err = unix.MoveMount(mfd, "", dfd, mountname, unix.MOVE_MOUNT_F_EMPTY_PATH)
	if err != nil {
		return -1, util.StatusWrapf(err, "Movemount mountname '%s' in file descriptor %d", mountname, dfd)
	}

	return mfd, nil
}

func (d *localDirectory) Mount(mountpoint path.Component, source string, fstype string) error {
	mfd, err := mountat(d.fd, fstype, source, mountpoint.String())
	if err != nil {
		return util.StatusWrap(err, "Mountat")
	}
	defer unix.Close(mfd)

	return nil
}
