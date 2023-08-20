package utils

import "github.com/gookit/color"

var (
	// colors
	Blue     = color.Blue.Render
	LBlue    = color.LightBlue.Render
	DBlue    = color.Blue.Darken().Render
	BlueCol  = color.Blue
	LBlueCol = color.LightBlue

	Cyan     = color.Cyan.Render
	LCyan    = color.LightCyan.Render
	CyanCol  = color.Cyan
	LCyanCol = color.LightCyan

	Magenta     = color.Magenta.Render
	LMagenta    = color.LightMagenta.Render
	MagentaCol  = color.Magenta
	LMagentaCol = color.LightMagenta

	Green     = color.Green.Render
	LGreen    = color.LightGreen.Render
	GreenCol  = color.Green
	LGreenCol = color.LightGreen

	Red     = color.Red.Render
	LRed    = color.LightRed.Render
	RedCol  = color.Red
	LRedCol = color.LightRed

	Yellow     = color.Yellow.Render
	LYellow    = color.LightYellow.Render
	YellowCol  = color.Yellow
	LYellowCol = color.LightYellow

	Gray     = color.FgDarkGray.Render
	LGray    = color.Gray.Render
	GrayCol  = color.FgDarkGray
	LGrayCol = color.Gray

	Normal = color.Normal.Render
)
