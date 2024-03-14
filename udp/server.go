package udp

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"wizard-beast-server-go/player"
	"wizard-beast-server-go/udp/request"
)

type Server struct {
	PlayerRepository player.Repository
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

	switch event {
	case "register":
		request.ProcessRegistration(
			id,
			payload,
			addr,
			server.PlayerRepository,
			client,
		)
	case "deregister":
		request.ProcessDeregistration(payload, server.PlayerRepository)
	case "update":
		request.ProcessAction(payload, server.PlayerRepository)
	case "acknowledge":
		request.ProcessAcknowledge(payload, server.PlayerRepository)
	default:
		client.Send("invalid", "unknown request", id)
	}
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
