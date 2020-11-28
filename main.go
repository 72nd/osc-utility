package main

import (
	"os"

	"github.com/72nd/osc-utility/src"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "osc-utility"
	app.Usage = "collection of tools for OSC"
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
					Value: "localhost",
				},
				cli.IntFlag{
					Name:  "port, p",
					Usage: "port of the OSC server",
					Value: 9000,
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
					Name:  "float, f",
					Usage: "float 32 argument (separate multiple by comma)",
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
				cli.StringFlag{
					Name:  "host",
					Usage: "host to run the server on",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "port, p",
					Usage: "port number to run the server on",
					Value: 9000,
				},
			},
			Action: serverAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func messageAction(c *cli.Context) error {
	msg := oscutility.Message{}
	if c.String("host") == "localhost" {
		logrus.Info("using default host (localhost)")
	}
	msg.Host = c.String("host")

	if c.Int("port") == 9000 {
		logrus.Info("using default port (9000)")
	}
	msg.Port = c.Int("port")

	if c.String("address") == "" {
		logrus.Error("no address specified (--address)")
		return nil
	}
	msg.Address = c.String("address")
	msg.SetBooleans(c.String("bool"))
	msg.SetStrings(c.String("string"))
	msg.SetIntegers(c.String("int"))
	msg.SetFloats(c.String("float"))
	msg.Send()
	return nil
}

func serverAction(c *cli.Context) error {
	srv := oscutility.Server{}
	srv.Host = c.String("host")
	if srv.Host == "127.0.0.1" {
		logrus.Info("using default host (127.0.0.1)")
	}
	srv.Port = c.Int("port")
	if c.Int("port") == 9000 {
		logrus.Info("using default port (9000)")
	}
	srv.Serve()
	return nil
}
