//go:build darwin || freebsd || windows
// +build darwin freebsd windows

package filesystem

func (d *localDirectory) Mount(mountpoint path.Component, source string, fstype string) error {
	return status.Error(codes.Unimplemented, "Mount is not supported")
}

func (d *localDirectory) Unmount(mountpoint path.Component) error {
	return status.Error(codes.Unimplemented, "Unmount is not supported")
}
