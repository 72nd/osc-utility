package oscutility

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/hypebeast/go-osc/osc"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Serve() {
	d := osc.NewStandardDispatcher()
	if err := d.AddMsgHandler("*", serverHandler); err != nil {
		slog.Error(err.Error())
	}
	srv := &osc.Server{
		Addr:       fmt.Sprintf("%s:%d", s.Host, s.Port),
		Dispatcher: d,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			slog.Error(err.Error())
		}
	}()
	slog.Info(fmt.Sprintf("OSC server runs on %s with port %d, Q + <Enter> to exit", s.Host, s.Port))
	promptForExit()
}

func serverHandler(msg *osc.Message) {
	slog.Debug("message received", "address", msg.Address, "arguments", msg.Arguments)
	attrs := []slog.Attr{
		slog.String("address", msg.Address),
	}
	var booleans []bool
	var strings []string
	var integers []int32
	var doubles []int64
	var floats []float32
	for _, arg := range msg.Arguments {
		switch arg := arg.(type) {
		case bool:
			booleans = append(booleans, arg)
		case string:
			strings = append(strings, arg)
		case int32:
			integers = append(integers, arg)
		case int64:
			doubles = append(doubles, arg)
		case float32:
			floats = append(floats, arg)
		}
	}
	attrs = appendSlogAttrIfNotEmpty(attrs, "booleans", booleans)
	attrs = appendSlogAttrIfNotEmpty(attrs, "strings", strings)
	attrs = appendSlogAttrIfNotEmpty(attrs, "integers", integers)
	attrs = appendSlogAttrIfNotEmpty(attrs, "doubles", doubles)
	attrs = appendSlogAttrIfNotEmpty(attrs, "floats", floats)
	slog.LogAttrs(context.Background(), slog.LevelInfo, "new message", attrs...)
}

func promptForExit() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		if input == "Q\n" {
			os.Exit(0)
		} else {
			fmt.Println("Q + <Enter> to exit")
		}
	}
}

func appendSlogAttrIfNotEmpty[T any](attrs []slog.Attr, key string, value []T) []slog.Attr {
	if len(value) == 0 {
		return attrs
	}
	return append(attrs, slog.Any(key, value))
}
