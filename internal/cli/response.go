package cli

import (
	"encoding/json"
	"fmt"
)

type AdvisorResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
	Data    struct {
		Temperature   int     `json:"temperature"`
		WindDirection string  `json:"wind_direction"`
		WindVelocity  float64 `json:"wind_velocity"`
		Humidity      float64 `json:"humidity"`
		Condition     string  `json:"condition"`
		Pressure      float64 `json:"pressure"`
		Icon          string  `json:"icon"`
		Sensation     int     `json:"sensation"`
		Date          string  `json:"date"`
	} `json:"data"`
}

type Conditions struct {
	City string
	Temp int
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp AdvisorResponse

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{},
			fmt.Errorf("Invalid Api response %s: %w", data, err)
	}

	conditions := Conditions{
		City: resp.Name,
		Temp: resp.Data.Temperature,
	}

	return conditions, nil
}
