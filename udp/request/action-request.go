package request

import "wizard-beast-server-go/entity"

type actionRequest struct {
	ID string `json:"id"`
}

func ProcessAction(
	data string,
	repository entity.PlayerRepository,
) {

}
