package internal

import "time"

type Candle struct {
	Timestamp time.Time `json:"timestamp"`
	Opening   float64   `json:"opening"`
	Closing   float64   `json:"closing"`
	Lowest    float64   `json:"lowest"`
	Highest   float64   `json:"highest"`
}
