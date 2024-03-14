package request

import "wizard-beast-server-go/entity"

type acknowledgeRequest struct {
	ID string `json:"id"`
}

func ProcessAcknowledge(
	data string,
	repository entity.PlayerRepository,
) {

}
