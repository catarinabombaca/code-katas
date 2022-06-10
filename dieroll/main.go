package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	Execute()
}

func Execute() {
	frequencies := roll(6, 1000000)
	fmt.Println(frequencies)
}

func roll(sidesOfDice int, rollTimes int) Frequency {
	frequencies := make(Frequency)

	for i := 0; i < rollTimes; i++ {
		rand.Seed(time.Now().UnixNano())
		frequencies[rand.Intn(sidesOfDice)+1]++
	}

	return frequencies
}

type Frequency map[int]int