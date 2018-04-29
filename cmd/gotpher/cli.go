package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/imdario/gotpher"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

var commands = []cli.Command{}

// newCliApp creates a new Gotpher CLI application.
func newCliApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "A little gopher companion for andOTP"
	app.Version = gotpher.Version
	app.Commands = commands
	return app
}

func loadFile(filename string) ([]byte, error) {
	if filename == "" {
		return nil, fmt.Errorf("input file not provided")
	}
	return ioutil.ReadFile(filename)
}

func getPassword(confirm bool) (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter your password",
		Mask:  '*',
	}
	if confirm {
		first, err := prompt.Run()
		if err != nil {
			return "", err
		}
		prompt.Label = "Enter your password again"
		second, err := prompt.Run()
		if err != nil {
			return "", err
		}
		if strings.Compare(first, second) != 0 {
			return "", fmt.Errorf("passwords don't match")
		}
		return second, nil
	}
	return prompt.Run()
}

func init() {
	// This is a hack to ensure only un/encrypted data is printed to stdout.
	readline.Stdout = os.Stderr
}
