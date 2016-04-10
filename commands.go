package main

import (
	"bufio"
	"os/exec"
	"sync"
)

// Run binary with arguments, check for errors, log output
func Cmd(cmdName string, cmdArgs ...string) (err error) {
	var wg sync.WaitGroup
	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorf("error creating StdoutPipe for tinc cmd: %s", err)
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	wg.Add(1)
	go CmdLogOutput(scanner, &wg)

	err = cmd.Start()
	if err != nil {
		log.Errorf("error starting cmd %s: %s", cmd.Args, err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		log.Errorf("error waiting for cmd %s: %s", cmd.Args, err)
		return err
	}

	wg.Wait()
	return nil
}

// Log output inside a goroutine
func CmdLogOutput(scanner *bufio.Scanner, wg *sync.WaitGroup) {
	defer wg.Done()
	for scanner.Scan() {
		log.Info(scanner.Text())
	}
}
