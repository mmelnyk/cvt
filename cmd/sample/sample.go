package main

import (
	"fmt"

	"github.com/mmelnyk/cvt"
)

func main() {

	fmt.Println(cvt.BrWhiteFg, "Example Table", cvt.ResetColor)

	if !cvt.IsTerminal() {
		fmt.Println("Is terminal available?")
	}

	fmt.Println(cvt.BrYellowFg, "Bright Yellow", cvt.BrGreenFg, "Bright Green", cvt.BrBlueFg, "Bright Blue", cvt.ResetColor, cvt.BlackFg)
	fmt.Println(cvt.YellowBg, "       Yellow", cvt.GreenBg, "       Green", cvt.BlueBg, "       Blue", cvt.ResetColor, cvt.BrWhiteFg)
	fmt.Println(cvt.RedBg, "          Red", cvt.MagentaBg, "     Magenta", cvt.CyanBg, "       Cyan", cvt.ResetColor)

}
