package udp

import (
	"errors"
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
	buf := make([]byte, 512)
	n, addr, err := conn.ReadFromUDP(buf)

	if err != nil {
		slog.Error("failed to read from UDP socket", err)
		return
	}

	event, payload, id, err := extract(buf[:n])

	if err != nil {
		slog.Warn("unable to extract infromation from the request", "request", buf[:n], "error", err)
		return
	}

	client := Client{conn, addr}

	go server.Processor.Process(
		id,
		event,
		payload,
		client,
	)
}

func extract(buf []byte) (event string, payload string, id string, err error) {
	split := strings.Split(string(buf[0:]), "--")

	if len(split) != 3 {
		err = errors.New("invalid request: the request should have 3 parts delimited by '--'")
		return
	}

	event = split[0]
	payload = split[1]
	id = split[2]
	return
}
