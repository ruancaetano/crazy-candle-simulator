package internal

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Candle struct {
	Opening float64 `json:"opening"`
	Closing float64 `json:"closing"`
	Lowest  float64 `json:"lowest"`
	Highest float64 `json:"highest"`
}

type CandleGenerator struct {
	interval time.Duration
	maxValue int
	minValue int
}

func NewCandleGenerator(interval time.Duration, max int, min int) *CandleGenerator {
	return &CandleGenerator{
		interval: interval,
		maxValue: max,
		minValue: min,
	}
}

func (g *CandleGenerator) Start() {

	for {

		lowAndHighValues := []float64{
			g.generateRandomNumber(g.minValue, g.maxValue),
			g.generateRandomNumber(g.minValue, g.maxValue),
		}
		sort.Float64s(lowAndHighValues)

		openAndCloseValues := []float64{
			g.generateRandomNumber(int(lowAndHighValues[0]), int(lowAndHighValues[1])),
			g.generateRandomNumber(int(lowAndHighValues[0]), int(lowAndHighValues[1])),
		}

		candle := &Candle{
			Lowest:  lowAndHighValues[0],
			Highest: lowAndHighValues[1],
			Opening: openAndCloseValues[0],
			Closing: openAndCloseValues[1],
		}

		fmt.Printf("['%d', %d, %d, %d, %d],\n", time.Now().Unix(), int(candle.Opening), int(candle.Highest), int(candle.Lowest), int(candle.Closing))

		time.Sleep(g.interval)
	}
}

func (g *CandleGenerator) generateRandomNumber(min int, max int) float64 {
	return float64(min) + (rand.Float64() * float64(max-min))
}
