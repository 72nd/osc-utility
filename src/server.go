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
	if len(booleans) != 0 {
		attrs = append(attrs, slog.Any("booleans", booleans))
	}
	if len(strings) != 0 {
		attrs = append(attrs, slog.Any("strings", strings))
	}
	if len(integers) != 0 {
		attrs = append(attrs, slog.Any("integers", integers))
	}
	if len(doubles) != 0 {
		attrs = append(attrs, slog.Any("doubles", doubles))
	}
	if len(floats) != 0 {
		attrs = append(attrs, slog.Any("floats", floats))
	}
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
