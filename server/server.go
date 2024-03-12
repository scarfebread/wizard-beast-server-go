package server

import (
	"fmt"
	"log/slog"
	"net"
	"strings"
	"wizard-beast-server-go/entity"
	"wizard-beast-server-go/server/request"
)

type Server struct {
	PlayerRepository entity.PlayerRepository
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
			slog.Error("Failed to close connection", err)
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
		slog.Error("Failed to read from UDP socket", err)
		return
	}

	message := string(buf[0:])
	split := strings.Split(message, "--")

	if len(split) != 3 {
		slog.Warn(fmt.Sprintf("Invalid request - %s", message))
		return
	}

	eventType := split[0]
	payload := split[1]
	requestId := split[2]

	switch eventType {
	case "register":
		{
			var registerRequest request.RegisterRequest
			err := registerRequest.Deserialise(payload)

			if err != nil {
				slog.Warn(fmt.Sprintf("Cannot deserialise %s", payload))
				return
			}

			player := entity.Player{
				Name: registerRequest.Name,
				Addr: addr,
			}
			server.PlayerRepository.AddPlayer(player)

			serialisedPlayer, err := player.Serialise()

			if err != nil {
				slog.Warn(fmt.Sprintf("Cannot serialise player %s", player.Name))
				return
			}

			_, err = conn.WriteToUDP([]byte(fmt.Sprintf("%s--%s--%s", "", serialisedPlayer, requestId)), addr)

			if err != nil {
				slog.Error("Unable to send message back to the client", err)
			}
		}
	}
}
