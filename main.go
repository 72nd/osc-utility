package main

import (
	"context"
	"log"
	"log/slog"
	"net/mail"
	"os"

	oscutility "github.com/72nd/osc-utility/src"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:    "osc-utility",
		Usage:   "utlity for working with OSC",
		Version: "0.2.2",
		Authors: []any{
			mail.Address{Name: "72nd", Address: "msg@frg72.com"},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "enable debug logging",
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Bool("debug") {
				slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
					Level: slog.LevelDebug,
				})))
			}
			return ctx, nil
		},
		Commands: []*cli.Command{
			{
				Name:    "message",
				Aliases: []string{"msg", "m"},
				Usage:   "send a message to a OSC server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "host",
						Usage:    "host of the OSC server",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "port of the OSC server",
						Value:   9000,
					},
					&cli.StringFlag{
						Name:     "address",
						Aliases:  []string{"adr", "a"},
						Usage:    "address of the message",
						Required: true,
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
						Name:    "bool",
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
						Value:   9000,
					},
				},
				Action: serverAction,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func messageAction(ctx context.Context, cmd *cli.Command) error {
	msg := oscutility.Message{}
	if cmd.String("host") == "localhost" {
		slog.Info("using default host (localhost)")
	}
	msg.Host = cmd.String("host")

	msg.Port = cmd.Int("port")

	msg.Address = cmd.String("address")
	if cmd.IsSet("bool") {
		msg.SetBooleans(cmd.String("bool"))
	}
	if cmd.IsSet("string") {
		msg.SetStrings(cmd.String("string"))
	}
	if cmd.IsSet("int") {
		msg.SetIntegers(cmd.String("int"))
	}
	if cmd.IsSet("float") {
		msg.SetFloats(cmd.String("float"))
	}
	msg.Send()
	return nil
}

func serverAction(ctx context.Context, cmd *cli.Command) error {
	srv := oscutility.Server{}
	srv.Host = cmd.String("host")
	srv.Port = cmd.Int("port")
	srv.Serve()
	return nil
}
