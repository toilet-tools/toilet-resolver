package logger

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/toilet-tools/toilet-resolver/utils"
)

type Logger struct {
	defaultPrefix string
}

func New(defaultPrefix string) Logger {
	e := Logger{defaultPrefix}
	return e
}

func (x Logger) prefix(str, title string, col color.Color) string {
	finishedPrefix := ""
	if len(x.defaultPrefix) != 0 {
		if strings.Contains(x.defaultPrefix, "<TIME>") {
			hours, minutes, seconds := time.Now().Clock()
			x.defaultPrefix = strings.ReplaceAll(x.defaultPrefix, "<TIME>", fmt.Sprintf("%s:%s:%s", strconv.Itoa(hours), strconv.Itoa(minutes), strconv.Itoa(seconds)))
		}
		if strings.Contains(x.defaultPrefix, "|") {
			x.defaultPrefix = strings.ReplaceAll(x.defaultPrefix, "|", utils.Normal("|"))
		}
		finishedPrefix += col.Render(x.defaultPrefix) + " | "
	}

	if len(title) != 0 {
		finishedPrefix += fmt.Sprintf("[%s] ", col.Render(title))
	}

	finishedPrefix += fmt.Sprintf("%s > ", col.Render(str))

	return finishedPrefix
}

func (x Logger) Info(str string) {
	prefix := x.prefix("i", "INFO", utils.LBlueCol)
	fmt.Println(prefix + str)
}

func (x Logger) Warn(str string) {
	prefix := x.prefix("!", "WARN", utils.LYellowCol)
	fmt.Println(prefix + str)
}

func (x Logger) Error(str string) {
	prefix := x.prefix("!", "ERROR", utils.LRedCol)
	fmt.Println(prefix + str)
}

func (x Logger) Debug(str string) {
	prefix := x.prefix("~", "DEBUG", utils.GrayCol)
	fmt.Println(prefix + str)
}

func (x Logger) Success(str string) {
	prefix := x.prefix("✓", "SUCCESS", utils.LGreenCol)
	fmt.Println(prefix + str)
}

func (x Logger) Invalid(str string) {
	prefix := x.prefix("-", "INVALID", utils.LRedCol)
	fmt.Println(prefix + str)
}

func (x Logger) Plus(str, title string) {
	prefix := x.prefix("+", title, utils.LGreenCol)
	fmt.Println(prefix + str)
}

func (x Logger) Money(str, title string) {
	prefix := x.prefix("$", title, utils.LCyanCol)
	fmt.Println(prefix + str)
}

func (x Logger) Log(char, str, title string, col color.Color) {
	prefix := x.prefix(char, title, col)
	fmt.Println(prefix + str)
}
