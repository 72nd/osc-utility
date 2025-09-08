package oscutility

import (
	"log/slog"

	"github.com/hypebeast/go-osc/osc"
)

type Message struct {
	Host     string
	Port     int
	Address  string
	Strings  []string
	Integers []int32
	Floats   []float32
	Booleans []bool
}

func (m *Message) Send() {
	client := osc.NewClient(m.Host, m.Port)
	msg := osc.NewMessage(m.Address)
	for _, value := range m.Booleans {
		msg.Append(value)
	}
	for _, value := range m.Strings {
		msg.Append(value)
	}
	for _, value := range m.Integers {
		msg.Append(value)
	}
	for _, value := range m.Floats {
		msg.Append(value)
	}
	if err := client.Send(msg); err != nil {
		slog.Error(err.Error())
	}
	slog.Debug(
		"sent message",
		"address", m.Address,
		"host", m.Host,
		"port", m.Port,
		"booleans", m.Booleans,
		"strings", m.Strings,
		"integers", m.Integers,
		"floats", m.Floats,
	)
}
