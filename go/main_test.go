package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseDie(t *testing.T) {
	// Test cases
	testCases := []struct {
		die       string
		expectedX int
		expectedY int
	}{
		{"1d6", 1, 6},
		{"2d10", 2, 10},
		{"3d20", 3, 20},
		{"1d2", 1, 2},
		{"12", 0, 0},
	}

	// Run tests
	for _, tc := range testCases {
		x, y := parseDie(tc.die)
		if x != tc.expectedX || y != tc.expectedY {
			t.Errorf("parseDie(%s) = (%d, %d), expected (%d, %d)", tc.die, x, y, tc.expectedX, tc.expectedY)
		}
	}
}

func TestGetRollRequests(t *testing.T) {
	// Save original os.Args
	origArgs := os.Args

	// Test case
	os.Args = []string{"cmd", "2d10", "3d6"}

	// Call getRollRequests
	requests := getRollRequests()

	// Expected result
	expected := []RollRequest{
		{Name: "2d10", Num: 2, Sides: 10},
		{Name: "3d6", Num: 3, Sides: 6},
	}

	// Check result
	if !reflect.DeepEqual(requests, expected) {
		t.Errorf("getRollRequests() = %v, expected %v", requests, expected)
	}

	// Restore original os.Args
	os.Args = origArgs
}

func TestRollRequest_Roll(t *testing.T) {
	// Test case
	request := RollRequest{Name: "2d10", Num: 2, Sides: 10}

	// Call Roll
	results := request.Roll()

	// Check result
	if len(results) != 2 {
		t.Errorf("len(results) = %d, expected 2", len(results))
	}
	for _, result := range results {
		if result.Type != 10 || result.Val < 1 || result.Val > 10 {
			t.Errorf("result = %v, expected 1 <= result <= 10", result)
		}
	}
}

func Test_roll(t *testing.T) {
	// Test case
	requests := []RollRequest{
		{Name: "2d10", Num: 2, Sides: 10},
		{Name: "3d6", Num: 3, Sides: 6},
	}

	// Call roll
	results := roll(requests)

	// Check result
	if len(results) != 5 {
		t.Errorf("len(results) = %d, expected 5", len(results))
	}
	for _, result := range results {
		if result.Type != 10 && result.Type != 6 {
			t.Errorf("result = %v, expected 10 or 6", result)
		}
		if result.Val < 1 || result.Val > result.Type {
			t.Errorf("result = %v, expected 1 <= result <= %d", result, result.Type)
		}
	}
}

func Test_sum(t *testing.T) {
	// Test case
	results := []DieCast{
		{Type: 10, Val: 1},
		{Type: 10, Val: 2},
		{Type: 6, Val: 1},
		{Type: 6, Val: 2},
		{Type: 6, Val: 3},
	}

	// Call sum
	total := sum(results)

	// Check result
	if total != 9 {
		t.Errorf("total = %d, expected 9", total)
	}
}
