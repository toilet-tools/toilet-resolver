package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/nathan-fiscaletti/consolesize-go"
)

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Title(title string) {
	cmd := exec.Command("cmd", "/c", "title", title)
	cmd.Run()
}

func Pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func PrintCenter(str string) {
	cols, _ := consolesize.GetConsoleSize()
	fmt.Printf(fmt.Sprintf("%%-%ds", cols), fmt.Sprintf(fmt.Sprintf("%%%ds", (cols+len(str))/2), str))
}
