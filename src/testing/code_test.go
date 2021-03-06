/**
 *  _             _ _| |_
 * | |           | |_   _|
 * | |___  _   _ | | |_|
 * | '_  \| | | || | | |
 * | | | || |_| || | | |
 * |_| |_|\___,_||_| |_|
 *
 * (c) Huli Inc
 */
// Package linter contains the examples shown on Golang Costa Rica meetup
// https://www.meetup.com/Golang-Costa-Rica/events/270127485/
package testing_test

import (
	"io/ioutil"
	"testing"

	td "github.com/fahernandez/go-linter-test-error-handling/src/testing"
)

func TestSum(t *testing.T) {
	x, y := 5, 4
	t.Logf("Number used: %v , %v", x, y)
	result := td.Sum(x, y)
	if result != 9 {
		t.Errorf("Should be 9, and find: %v", result)
	}
}

var sumTest = []struct {
	firsNumber   int
	secondNumber int
	result       int
	message      string
}{
	{1, 1, 2, "Positive + Positive number ok"},
	{1, -1, 0, "Positive + Negative number ok"},
	{-1, -1, -2, "Negative + Negative number ok"},
}

func TestTableSum(t *testing.T) {

	for _, tt := range sumTest {
		t.Log(tt.message)
		result := td.Sum(tt.firsNumber, tt.secondNumber)
		if result != tt.result {
			t.Errorf("Fail, %v + %v = %v", tt.firsNumber, tt.secondNumber, tt.result)
		}
	}
}

func TestRead(t *testing.T) {
	t.Log("Reading test")
	data, err := ioutil.ReadFile("data/data.data")
	if err != nil {
		t.Fatal("Error trying to open the file")
	}
	if string(data) != "It should have this text" {
		t.Fatal("Info error")
	}
}
