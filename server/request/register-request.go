package request

import "encoding/json"

type RegisterRequest struct {
	Name string `json:"name"`
}

func (req *RegisterRequest) Deserialise(data string) error {
	var personData map[string]interface{}

	err := json.Unmarshal([]byte(data), &personData)

	if err != nil {
		return err
	}

	req.Name = personData["name"].(string)

	return nil
}
