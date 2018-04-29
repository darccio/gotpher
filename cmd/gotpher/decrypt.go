package main

import (
	"fmt"

	"github.com/imdario/gotpher"
	"github.com/urfave/cli"
)

func init() {
	decryptCmd := cli.Command{
		Name:    "decrypt",
		Aliases: []string{"d"},
		Usage:   "decrypt a backup file",
		Action:  decrypt,
	}
	commands = append(commands, decryptCmd)
}

func decrypt(ctx *cli.Context) error {
	input, err := loadFile(ctx.Args().First())
	if err != nil {
		return err
	}
	password, err := getPassword(false)
	if err != nil {
		return err
	}
	result, err := gotpher.Decrypt(input, password)
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	return nil
}
