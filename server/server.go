package server

import (
	"log/slog"
	"net"
)

func Start() error {
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

	defer conn.Close() // TODO unhandled error

	for {
		processEvent(conn)
	}
}

func processEvent(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		slog.Error("Failed to read from UDP socket", err)
		return
	}

	// string(buf[0:]) // TODO do something with this
	_, err = conn.WriteToUDP(buf[0:], addr)

	if err != nil {
		slog.Error("Unable to send message back to the client", err)
	}
}
