package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/gobuffalo/packr"
)

const scriptFolder = "/tmp/scripts"
const scriptName = "hello.sh"

func main() {

	box := packr.NewBox(scriptFolder)
	if !box.Has(scriptName) {
		log.Fatalf("Specified Script \"%s\" does not exist", scriptName)
	}
	if err := runScript(box, scriptName); err != nil {
		log.Fatalf("Error running %s", err)
	}
}

func runScript(box packr.Box, path string) error {

	tmpfile, _ := createTmpFile(box, path)
	defer os.Remove(tmpfile.Name()) // clean up

	cmd := exec.Command(tmpfile.Name())
	_, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)
	return nil

}

func createTmpFile(box packr.Box, script string) (*os.File, error) {
	content := box.Bytes(script)
	tmpfile, err := ioutil.TempFile("", script)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if err := os.Chmod(tmpfile.Name(), 0700); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpfile, nil

}
