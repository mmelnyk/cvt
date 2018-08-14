package cvt

var (
	isTerminalCached = isTerminal()
)

// IsTerminal is returning if StdOut is terminal
func IsTerminal() bool {
	return isTerminalCached
}
