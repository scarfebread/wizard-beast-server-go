package request

import (
	"net"
	"wizard-beast-server-go/player"
	"wizard-beast-server-go/udp"
)

type Processor struct {
	PlayerRepository player.Repository
}

func (processor Processor) Process(
	id string,
	event string,
	payload string,
	addr *net.UDPAddr,
	client udp.Client,
) {
	switch event {
	case "register":
		ProcessRegistration(
			id,
			payload,
			addr,
			processor.PlayerRepository,
			client,
		)
	case "deregister":
		ProcessDeregistration(payload, processor.PlayerRepository)
	case "update":
		ProcessAction(payload, processor.PlayerRepository)
	case "acknowledge":
		ProcessAcknowledge(payload, processor.PlayerRepository)
	default:
		client.Send("invalid", "unknown request", id)
	}
}
