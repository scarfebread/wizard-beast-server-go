package request

import "wizard-beast-server-go/entity"

type deregisterRequest struct {
	ID string `json:"id"`
}

func ProcessDeregistration(
	data string,
	repository entity.PlayerRepository,
) {

}
