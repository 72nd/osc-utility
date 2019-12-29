package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "prfm-osc"
	app.Usage = "some OSC tools for debugging and alike"
	app.Action = func(c *cli.Context) error {
		_ = cli.ShowCommandHelp(c, c.Command.Name)
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "message",
			Aliases: []string{"msg", "m"},
			Usage:   "send a message to a OSC server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Usage: "host of the OSC server",
				},
				cli.IntFlag{
					Name:  "port, p",
					Usage: "port of the OSC server",
				},
				cli.StringFlag{
					Name:  "address, adr, a",
					Usage: "address of the message",
				},
				cli.StringFlag{
					Name:  "string, str, s",
					Usage: "string argument (separate multiple by comma)",
				},
				cli.StringFlag{
					Name:  "int, i",
					Usage: "integer 32 argument (separate multiple by comma)",
				},
				cli.StringFlag{
					Name:  "bool, b",
					Usage: "boolean argument (separate multiple by comma)",
				},
			},
			Action: messageAction,
		},
		{
			Name:    "server",
			Aliases: []string{"srv", "s"},
			Usage:   "OSC server on localhost, prints incoming messages",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Usage: "port number to run the server on",
				},
			},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func messageAction(c *cli.Context) error {
	if len(c.Args()) == 0 {
		_ = cli.ShowCommandHelp(c, c.Command.Name)
		logrus.Error("need some more information")
	}
	if c.String("host") == "" {
		logrus.Error("no host specified (--host)")
		return nil
	}
	if c.Int("port") == 0 {
		logrus.Error("no port specified (--port)")
		return nil
	}
	if c.String("address") == "" {
		logrus.Error("no address specified (--address)")
		return nil
	}
	return nil
}
