// +build darwin dragonfly freebsd netbsd openbsd

package cvt

import (
	"syscall"
)

type Termios syscall.Termios

const ioctlReadTermios = syscall.TIOCGETA
