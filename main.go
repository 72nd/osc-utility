package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/mail"
	"os"
	"slices"
	"strings"

	oscutility "github.com/72nd/osc-utility/src"
	"github.com/urfave/cli/v3"
)

var (
	BoolTrue  = []string{"true", "t", "1"}
	BoolFalse = []string{"false", "f", "0"}
)

func main() {
	app := &cli.Command{
		Name:    "osc-utility",
		Usage:   "utlity for working with OSC",
		Version: "0.3.0",
		Authors: []any{
			mail.Address{Name: "72nd", Address: "msg@frg72.com"},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "enable debug logging",
			},
			&cli.BoolFlag{
				Name:  "json-log",
				Usage: "output logs in json format",
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Bool("debug") {
				slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
					Level: slog.LevelDebug,
				})))
			}
			if cmd.Bool("json-log") {
				slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
					Level: slog.LevelInfo,
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
						Name:  "host",
						Usage: "host of the OSC server",
					},
					&cli.IntFlag{
						Name:     "port",
						Aliases:  []string{"p"},
						Usage:    "port of the OSC server",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "address",
						Aliases:  []string{"adr", "a"},
						Usage:    "address of the message",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:    "string",
						Aliases: []string{"str", "s"},
						Usage:   "string argument(s)",
					},
					&cli.Int32SliceFlag{
						Name:    "int",
						Aliases: []string{"i"},
						Usage:   "integer 32 argument(s)",
					},
					&cli.Float32SliceFlag{
						Name:    "float",
						Aliases: []string{"f"},
						Usage:   "float 32 argument(s)",
					},
					&cli.StringSliceFlag{
						Name:    "bool",
						Aliases: []string{"b"},
						Usage:   "boolean argument(s)",
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
						Name:     "port",
						Aliases:  []string{"p"},
						Usage:    "port number to run the server on",
						Required: true,
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
	host := cmd.String("host")
	if !cmd.IsSet("host") {
		slog.Info("no host provided, using default (127.0.0.1)")
		host = "127.0.0.1"
	}

	booleans, err := parseBoolArg(cmd.StringSlice("bool"))
	if err != nil {
		return err
	}

	msg := oscutility.Message{
		Host:     host,
		Port:     cmd.Int("port"),
		Address:  cmd.String("address"),
		Strings:  cmd.StringSlice("string"),
		Integers: cmd.Int32Slice("int"),
		Floats:   cmd.Float32Slice("float"),
		Booleans: booleans,
	}

	msg.Send()
	return nil
}

func serverAction(ctx context.Context, cmd *cli.Command) error {
	srv := oscutility.Server{}
	srv.Host = cmd.String("host")
	srv.Port = cmd.Int("port")
	srv.Serve(!cmd.Bool("json-log"))
	return nil
}

func parseBoolArg(arg []string) ([]bool, error) {
	var rsl []bool
	for _, value := range arg {
		if slices.Contains(BoolTrue, value) {
			rsl = append(rsl, true)
			continue
		}
		if slices.Contains(BoolFalse, value) {
			rsl = append(rsl, false)
			continue
		}
		return nil, fmt.Errorf("invalid boolean value: %s, use one of [%s] for true or [%s] for false", value, strings.Join(BoolTrue, ", "), strings.Join(BoolFalse, ", "))
	}
	return rsl, nil
}
