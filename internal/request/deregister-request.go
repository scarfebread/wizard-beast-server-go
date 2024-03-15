package request

import "github.com/scarfebread/wizard-beast-server-go/internal/player"

type deregisterRequest struct {
	ID string `json:"id"`
}

func ProcessDeregistration(
	data string,
	repository player.Repository,
) {

}
