package request

import "wizard-beast-server-go/player"

type acknowledgeRequest struct {
	ID string `json:"id"`
}

func ProcessAcknowledge(
	data string,
	repository player.Repository,
) {

}
