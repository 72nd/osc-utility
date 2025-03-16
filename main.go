package main

import (
	"os"

	oscutility "github.com/72nd/osc-utility/src"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "osc-utility",
		Usage:   "utlity for working with OSC",
		Version: "0.2.1",
		Authors: []*cli.Author{
			{
				Name:  "72nd",
				Email: "msg@frg72.com",
			},
		},
		Action: func(c *cli.Context) error {
			_ = cli.ShowCommandHelp(c, c.Command.Name)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "message",
				Aliases: []string{"msg", "m"},
				Usage:   "send a message to a OSC server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "host of the OSC server",
						Value: "localhost",
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "port of the OSC server",
						Value:   9000,
					},
					&cli.StringFlag{
						Name:    "address",
						Aliases: []string{"adr", "a"},
						Usage:   "address of the message",
					},
					&cli.StringFlag{
						Name:    "string",
						Aliases: []string{"str", "s"},
						Usage:   "string argument (separate multiple values by comma)",
					},
					&cli.IntFlag{
						Name:    "int",
						Aliases: []string{"i"},
						Usage:   "integer 32 argument (separate multiple values by comma)",
					},
					&cli.StringFlag{
						Name:    "float",
						Aliases: []string{"f"},
						Usage:   "float 32 argument (separate multiple values by comma)",
					},
					&cli.StringFlag{
						Name:    "bool, b",
						Aliases: []string{"b"},
						Usage:   "boolean argument (separate multiple values by comma)",
					},
				},
				Action: messageAction,
			},
			{
				Name:    "server",
				Aliases: []string{"srv", "s"},
				Usage:   "OSC server on localhost, prints incoming messages",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Usage: "host to run the server on",
						Value: "127.0.0.1",
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "port number to run the server on",
					},
				},
				Action: serverAction,
			},
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

	msg.Port = c.Int("port")

	if c.Int("port") == 0 {
		logrus.Error("no port specified (--port)")
		return nil
	}
	if c.String("address") == "" {
		logrus.Error("no address specified (--address)")
		return nil
	}
	msg.Address = c.String("address")
	if c.IsSet("bool") {
		msg.SetBooleans(c.String("bool"))
	}
	if c.IsSet("string") {
		msg.SetStrings(c.String("string"))
	}
	if c.IsSet("int") {
		msg.SetIntegers(c.String("int"))
	}
	if c.IsSet("float") {
		msg.SetFloats(c.String("float"))
	}
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
	if c.Int("port") == 0 {
		logrus.Error("no port specified (--port)")
		return nil
	}
	srv.Serve()
	return nil
}
