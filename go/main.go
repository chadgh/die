package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

// DieCast represents the result of a single die roll.
type DieCast struct {
	Type int
	Val  int
}

// RollRequest represents a request to roll a number of dice with a given number of sides.
type RollRequest struct {
	Name  string
	Num   int
	Sides int
}

// Roll the dice the specified number of times and return the results.
func (r *RollRequest) Roll() []DieCast {
	results := make([]DieCast, r.Num)
	for i := 0; i < r.Num; i++ {
		results[i] = DieCast{
			Type: r.Sides,
			Val:  rand.Intn(r.Sides) + 1,
		}
	}
	return results
}

func getRollRequests() []RollRequest {
	flag.Parse()
	args := flag.Args()
	requests := make([]RollRequest, len(args))
	for i, arg := range args {
		num, sides := parseDie(arg)
		requests[i] = RollRequest{
			Name:  arg,
			Num:   num,
			Sides: sides,
		}
	}
	return requests
}

func parseDie(die string) (int, int) {
	var num, sides int
	_, err := fmt.Sscanf(die, "%dd%d", &num, &sides)
	if err != nil {
		fmt.Printf("Usage: go run %s NdM [NdM ...]\n", os.Args[0])
		return 0, 0
	}
	return num, sides
}

// Roll the dice and return the results.
func roll(requests []RollRequest) []DieCast {
	results := make([]DieCast, 0)
	for _, req := range requests {
		results = append(results, req.Roll()...)
	}
	return results
}

func printResults(results []DieCast, total int) {
	fmt.Printf("Total: %d\n", total)
	for _, result := range results {
		fmt.Printf("   d%d: %d\n", result.Type, result.Val)
	}
}

func sum(results []DieCast) int {
	total := 0
	for _, r := range results {
		total += r.Val
	}
	return total
}

func main() {
	// Parse the command line arguments to get the roll requests.
	rollRequests := getRollRequests()

	// Roll the dice and get the results.
	results := roll(rollRequests)
	// Calculate the total of all the dice.
	total := sum(results)

	// Print the results.
	printResults(results, total)
}
