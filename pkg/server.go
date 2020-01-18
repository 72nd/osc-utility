package pkg

import (
	"bufio"
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"github.com/sirupsen/logrus"
	"os"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Serve() {
	d := osc.NewStandardDispatcher()
	if err := d.AddMsgHandler("*", serverHandler); err != nil {
		logrus.Error(err)
	}
	srv := &osc.Server{
		Addr:       fmt.Sprintf("%s:%d", s.Host, s.Port),
		Dispatcher: d,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()
	logrus.Infof("OSC server runs on %s with port %d, Q + <Enter> to exit", s.Host, s.Port)
	promptForExit()
}

func serverHandler(msg *osc.Message) {
	fields := logrus.Fields{
		"address": msg.Address,
	}
	var booleans []bool
	var strings []string
	var integers []int32
	var doubles []int64
	var floats []float32
	for _, arg := range msg.Arguments {
		switch arg.(type) {
		case bool:
			booleans = append(booleans, arg.(bool))
		case string:
			strings = append(strings, arg.(string))
		case int32:
			integers = append(integers, arg.(int32))
		case int64:
			doubles = append(doubles, arg.(int64))
		case float32:
			floats = append(floats, arg.(float32))
		}
	}
	if len(booleans) != 0 {
		fields["booleans"] = booleans
	}
	if len(strings) != 0 {

		fields["strings"] = strings
	}
	if len(integers) != 0 {
		fields["integers"] = integers
	}
	if len(doubles) != 0 {
		fields["doubles"] = doubles
	}
	if len(floats) != 0 {
		fields["floats"] = floats
	}
	logrus.WithFields(fields).Info("new message")
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
