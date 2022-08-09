package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://127.0.0.1:3541/get?id=test3",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
