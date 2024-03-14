package player

import (
	"encoding/json"
	"net"
)

type Player struct {
	Id   string
	Name string
	X    float32
	Y    float32
	Addr *net.UDPAddr
}

func (p Player) Serialise() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   p.Id,
		"name": p.Name,
		"x":    p.X,
		"y":    p.Y,
	})
}
