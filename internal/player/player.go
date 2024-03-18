package player

import (
	"encoding/json"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
)

type Player struct {
	ID         string
	Name       string
	X          float32
	Y          float32
	Input      Input
	InputQueue []Input
	Client     udp.Client
}

type Action struct {
	Key     int  `json:"key"`
	Pressed bool `json:"currentlyPressed"`
}

type Input struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

func (p Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   p.ID,
		"name": p.Name,
		"x":    p.X,
		"y":    p.Y,
	})
}
