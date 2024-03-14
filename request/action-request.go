package request

import "wizard-beast-server-go/player"

type actionRequest struct {
	ID string `json:"id"`
}

func ProcessAction(
	data string,
	repository player.Repository,
) {

}
