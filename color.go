package cvt

// See ANSI Escape Codes - https://en.wikipedia.org/wiki/ANSI_escape_code

const (
	// ESCape codes
	ESC = "\033"
	CSI = "\033["

	//Colors
	ResetColor = "\033[0m"
	Bold       = "\033[1m"
	Faint      = "\033[2m"
	Normal     = "\033[22m"

	BlackFg   = "\033[30m"
	RedFg     = "\033[31m"
	GreenFg   = "\033[32m"
	YellowFg  = "\033[33m"
	BlueFg    = "\033[34m"
	MagentaFg = "\033[35m"
	CyanFg    = "\033[36m"
	WhiteFg   = "\033[37m"

	BlackBg   = "\033[40m"
	RedBg     = "\033[41m"
	GreenBg   = "\033[42m"
	YellowBg  = "\033[43m"
	BlueBg    = "\033[44m"
	MagentaBg = "\033[45m"
	CyanBg    = "\033[46m"
	WhiteBg   = "\033[47m"

	BrBlackFg   = "\033[90m"
	BrRedFg     = "\033[91m"
	BrGreenFg   = "\033[92m"
	BrYellowFg  = "\033[93m"
	BrBlueFg    = "\033[94m"
	BrMagentaFg = "\033[95m"
	BrCyanFg    = "\033[96m"
	BrWhiteFg   = "\033[97m"

	BrBlackBg   = "\033[100m"
	BrRedBg     = "\033[101m"
	BrGreenBg   = "\033[102m"
	BrYellowBg  = "\033[103m"
	BrBlueBg    = "\033[104m"
	BrMagentaBg = "\033[105m"
	BrCyanBg    = "\033[106m"
	BrWhiteBg   = "\033[107m"

	// Controls
	EraseDisplay = "\033[2J"
	EraseLine    = "\033[K"
	MoveTop      = "\033[H" // Format is CSI Line;Column H
	MoveBegin    = "\033[G" // Format is CSI Column G
	SavePos      = "\033[s"
	RestorePos   = "\033[u"
)
