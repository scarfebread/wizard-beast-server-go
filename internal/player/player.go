package player

import (
	"encoding/json"
	"github.com/scarfebread/wizard-beast-server-go/internal/udp"
)

type Player struct {
	Id     string
	Name   string
	X      float32
	Y      float32
	Client udp.Client
}

func (p Player) Serialise() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   p.Id,
		"name": p.Name,
		"x":    p.X,
		"y":    p.Y,
	})
}
