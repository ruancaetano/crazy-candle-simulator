package main

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/generator/internal"
	"time"
)

func main() {

	generator := internal.NewCandleGenerator(time.Second, 0, 1000)

	generator.Start()
}
