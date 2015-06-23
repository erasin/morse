package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/erasin/morse"
)

func main() {
	app := cli.NewApp()
	app.Name = "morse-go"
	app.Usage = "morese decode|encode string"

	mo := morse.NewMorse()

	var decode = &cli.Command{
		Name:      "decode",
		ShortName: "de",
		Usage:     "decode string",
		Action: func(c *cli.Context) {
			fmt.Println(mo.Decode(c.Args().First()))
		},
	}

	var encode = &cli.Command{
		Name:      "encode",
		ShortName: "en",
		Usage:     "encode string",
		Action: func(c *cli.Context) {
			fmt.Println(mo.Encode(c.Args().First()))
		},
	}

	app.Commands = []cli.Command{*decode, *encode}
	app.Run(os.Args)

}
