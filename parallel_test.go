package fanout

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/joeshaw/gengen/generic"
)

func TestParallelRun(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	inputNum := 50

	vals := []generic.T{}
	for i := 0; i < inputNum; i++ {
		vals = append(vals, inputVal{number: i, giveError: ""})
	}

	results, err := ParallelRun(100, w, vals)
	if err != nil {
		t.Error(err)
	}
	if len(results) != inputNum {
		t.Errorf("results count is not %d, is %d", inputNum, len(results))
	}
}

func TestParallelRunWithError(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	inputNum := 20

	vals := []generic.T{}
	for i := 0; i < inputNum; i++ {
		vals = append(vals, inputVal{number: i, giveError: ""})
	}

	_, err := ParallelRun(20, errw, vals)

	// fmt.Println(err)
	if err == nil {
		t.Error("should have error")
	}
}

func errw(i generic.T) (r generic.T, e error) {
	e = errors.New("I am an error")
	return
}

type inputVal struct {
	giveError string
	number    int
}

func w(i generic.T) (r generic.T, e error) {
	var val = i.(inputVal)
	result := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(result))
	fmt.Println(val.number, " [", result, "ms ]")
	r = result
	if len(val.giveError) > 0 {
		e = errors.New(val.giveError)
	}
	return
}
