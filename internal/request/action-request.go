package request

import "github.com/scarfebread/wizard-beast-server-go/internal/player"

type actionRequest struct {
	ID string `json:"id"`
}

func ProcessAction(
	data string,
	repository player.Repository,
) {

}
