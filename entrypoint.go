package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func queryEscape(s string) string {
	schemePath := strings.Split(s, "://")
	if len(schemePath) != 2 {
		panic("invalid database uri")
	}
	usernamePassword := strings.Split(schemePath[1], "@")
	password := strings.Split(usernamePassword[0], ":")
	if len(password) == 2 {
		s = strings.Replace(s, password[1], url.QueryEscape(password[1]), 3)
	}
	return s
}

func main() {
	args := os.Args
	migCmd := "migrate"

	scheme := strings.Split(args[2], "://")[0]
	if scheme == "sqlite3" {
		migCmd = "/sqlite-migrate"
	}

	arguments := []string{
		"-path",
		args[1],
		"-database",
		queryEscape(args[2]),
		"-prefetch",
		args[3],
		"-lock-timeout",
		args[4],
	}

	if len(args[5]) > 0 {
		arguments = append(arguments, "-verbose")
	}

	if len(args[6]) > 0 {
		arguments = append(arguments, "-version")
	}

	arguments = append(arguments, args[7])

	cmd := exec.Command(migCmd, arguments...)
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Printf("error: %v", err)
		var exitError *exec.ExitError
		ok := errors.As(err, &exitError)
		if ok {
			os.Exit(exitError.ExitCode())
		}
	}

	log.Println(stdBuffer.String())
}
