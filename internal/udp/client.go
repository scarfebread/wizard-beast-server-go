package udp

import (
	"fmt"
	"log/slog"
	"net"
)

type Client struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

func (client Client) Send(event, message, id string) {
	_, err := client.conn.WriteToUDP([]byte(fmt.Sprintf("%s--%s--%s", event, message, id)), client.addr)

	if err != nil {
		slog.Error(fmt.Sprintf("unable to send message back to the client %s", client.addr.IP), err)
	}
}
