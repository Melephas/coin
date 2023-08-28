package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var coinStates = []string{
	"Heads",
	"Tails",
}

var decisionStates = []string{
	"Yes",
	"No",
}

var trialsFlag int
var nFlag int
var decisionFlag bool

func init() {
	flag.IntVar(&trialsFlag, "t", 1, "The number of trials (coin flips, decisions, or random numbers)")
	flag.IntVar(&nFlag, "n", 0, "Pick a number between 1 and the provided value")
	flag.BoolVar(&decisionFlag, "d", false, "Choose between yes and no instead of heads and tails")

	log.SetPrefix("coin: ")
	log.SetFlags(0)

	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	if nFlag > 0 && decisionFlag {
		log.Fatalln("Can't decide between yes/no and pick from range at the same time, use either -d or -n<val>")
	}

	switch {
	case decisionFlag:
		for i := 0; i < trialsFlag; i++ {
			fmt.Println(decisionStates[rand.Intn(len(decisionStates))])
		}
	case nFlag > 0:
		for i := 0; i < trialsFlag; i++ {
			fmt.Println(rand.Intn(nFlag) + 1)
		}
	default:
		for i := 0; i < trialsFlag; i++ {
			fmt.Println(coinStates[rand.Intn(len(coinStates))])
		}
	}

	os.Exit(0)
}
