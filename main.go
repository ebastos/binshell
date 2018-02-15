package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/gobuffalo/packr"
)

const scriptFolder = "/tmp/scripts"
const scriptName = "hello.sh"

func main() {

	box := packr.NewBox(scriptFolder)
	if !box.Has(scriptName) {
		log.Fatalf("Specified Script \"%s\" does not exist", scriptName)
	}
	runScript(box, scriptName)
}

func runScript(box packr.Box, script string) error {

	s := box.String(scriptName)

	sheBang := strings.SplitAfter(s, "\n")[0]
	interpreter := strings.TrimSpace(strings.SplitAfter(sheBang, "!")[1])

	cmd := exec.Command(interpreter)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, s)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)
	return nil

}
