package oscutility

import (
	"github.com/hypebeast/go-osc/osc"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
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

func (m *Message) SetStrings(input string) {
	if input == "" {
		return
	}
	parts := strings.Split(input, ",")
	m.Strings = parts
}

func (m *Message) SetIntegers(input string) {
	parts := strings.Split(input, ",")
	if len(parts) == 1 && parts[0] == "" {
		return
	}
	for _, part := range parts {
		value, err := strconv.ParseInt(part, 0, 32)
		if err != nil {
			logrus.Warnf("argument %s could not be parsed as a int32, ignoring this one", part)
			continue
		}
		m.Integers = append(m.Integers, int32(value))
	}
}

func (m *Message) SetFloats(input string) {
	parts := strings.Split(input, ",")
	if len(parts) == 1 && parts[0] == "" {
		return
	}
	for _, part := range parts {
		value, err := strconv.ParseFloat(part, 32)
		if err != nil {
			logrus.Warnf("argument %s could not be parsed as a float, ignoring this one", part)
			continue
		}
		m.Floats = append(m.Floats, float32(value))
	}
}

func (m *Message) SetBooleans(input string) {
	parts := strings.Split(input, ",")
	if len(parts) == 1 && parts[0] == "" {
		return
	}
	for _, part := range parts {
		var value bool
		if part == "true" || part == "t" || part == "1" {
			value = true
		} else if part == "false" || part == "f" || part == "0" {
			value = false
		} else {
			logrus.Warnf("argument %s could not be parsed as boolean, ignoring this one", part)
			continue
		}
		m.Booleans = append(m.Booleans, value)
	}
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
		logrus.Error(err)
	}
}
