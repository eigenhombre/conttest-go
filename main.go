package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func bashCommand(args []string) ([]byte, error) {
	shellCommand := strings.Join(args, " ")
	cmdResult := exec.Command("bash", "-c", shellCommand)
	return cmdResult.CombinedOutput() // ([]byte, error)
}

func main() {
	runCmd := func() {
		out, _ := bashCommand(os.Args[1:])
		fmt.Printf("%s", out)
	}

	lastTime := time.Now()
	runCmd()
	for {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		e := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			thisTime := info.ModTime()
			if thisTime.After(lastTime) {
				lastTime = thisTime
				runCmd()
			}
			return nil
		})
		if e != nil {
			log.Fatal(e)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
