package request

import "github.com/scarfebread/wizard-beast-server-go/internal/player"

type acknowledgeRequest struct {
	ID string `json:"id"`
}

func ProcessAcknowledge(
	data string,
	repository player.Repository,
) {

}
