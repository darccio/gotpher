package main

import (
	"os"

	"github.com/imdario/gotpher"
	"github.com/urfave/cli"
)

func init() {
	encryptCmd := cli.Command{
		Name:    "encrypt",
		Aliases: []string{"e"},
		Usage:   "encrypt a backup file",
		Action:  encrypt,
	}
	commands = append(commands, encryptCmd)
}

func encrypt(ctx *cli.Context) error {
	input, err := loadFile(ctx.Args().First())
	if err != nil {
		return err
	}
	password, err := getPassword(true)
	if err != nil {
		return err
	}
	result, err := gotpher.Encrypt(input, password)
	if err != nil {
		return err
	}
	os.Stdout.Write(result)
	return nil
}
