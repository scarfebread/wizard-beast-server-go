package udp

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"
)

type Server struct {
	Processor Processor
}

type Processor interface {
	Process(
		id string,
		event string,
		payload string,
		addr *net.UDPAddr,
		client Client,
	)
}

func (server Server) Start() error {
	address, err := net.ResolveUDPAddr("udp", ":9002")

	if err != nil {
		return err
	}

	conn, err := net.ListenUDP(
		"udp",
		address,
	)

	if err != nil {
		return err
	}

	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			slog.Error("failed to close connection", err)
		}
	}(conn)

	for {
		server.processEvent(conn)
	}
}

func (server Server) processEvent(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		slog.Error("failed to read from UDP socket", err)
		return
	}

	event, payload, id, err := extract(buf)

	if err != nil {
		slog.Warn(fmt.Sprintf("invalid request - %s", string(buf[0:])), err)
		return
	}

	client := Client{conn, addr}

	server.Processor.Process(
		id,
		event,
		payload,
		addr,
		client,
	)
}

func extract(buf [512]byte) (event string, payload string, id string, err error) {
	split := strings.Split(string(buf[0:]), "--")

	if len(split) != 3 {
		err = errors.New("invalid request")
		return
	}

	event = split[0]
	payload = split[1]
	id = split[2]
	return
}
