package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func queryEscape(s string) string {
	split := strings.Split(s, "://")
	usernamePassword := strings.Split(split[1], "@")
	password := strings.Split(usernamePassword[0], ":")
	replaced := strings.Replace(s, password[1], url.QueryEscape(password[1]), 3)

	return replaced
}

func main() {
	args := os.Args

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

	cmd := exec.Command("migrate", arguments...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
	}

	fmt.Printf("Result: \n%s\n", string(out))
}
