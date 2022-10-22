package entities

import "time"

type Candle struct {
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Opening   float64   `json:"opening" bson:"opening"`
	Closing   float64   `json:"closing" bson:"closing"`
	Lowest    float64   `json:"lowest" bson:"lowest"`
	Highest   float64   `json:"highest" bson:"highest"`
}
