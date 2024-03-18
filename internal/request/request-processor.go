package request

import (
	"github.com/scarfebread/wizard-beast-server-go/internal/game"
	"github.com/scarfebread/wizard-beast-server-go/internal/player"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
)

type Processor struct {
	PlayerRepository player.Repository
	Simulator        game.Simulator
}

func (processor Processor) Process(
	id string,
	event string,
	payload string,
	client udp.Client,
) {
	switch event {
	case "register":
		ProcessRegistration(
			id,
			payload,
			processor.PlayerRepository,
			client,
		)
	case "deregister":
		ProcessDeregistration(
			payload,
			processor.PlayerRepository,
			client,
		)
	case "update":
		ProcessAction(payload, processor.Simulator)
	case "acknowledge":
		ProcessAcknowledge(payload, processor.PlayerRepository)
	default:
		client.Send("invalid", "unknown request", id)
	}
}
