package request

import "wizard-beast-server-go/player"

type deregisterRequest struct {
	ID string `json:"id"`
}

func ProcessDeregistration(
	data string,
	repository player.Repository,
) {

}
