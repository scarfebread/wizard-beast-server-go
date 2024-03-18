package udp

import (
	"fmt"
	"log/slog"
	"net"
)

type Client struct {
	conn *net.UDPConn
	Addr *net.UDPAddr
}

func (c Client) Send(event, message, id string) {
	_, err := c.conn.WriteToUDP([]byte(fmt.Sprintf("%s--%s--%s", event, message, id)), c.Addr)

	if err != nil {
		slog.Error("unable to connect to client", "client", c.Addr.IP, "error", err)
	}
}

func (c Client) Validate(c2 Client) bool {
	return c.Addr.IP.String() == c2.Addr.IP.String()
}
