package publishable

const (
	Connect    = "connect"
	Disconnect = "disconnect"
	Move       = "move"
)

type PlayerAction struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Action string  `json:"action"`
}
