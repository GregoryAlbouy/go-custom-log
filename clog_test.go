package clog

import (
	"errors"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	runError()
	runSuccess()
	Printlb([]int{1, 2, 3}, Blue("Some numbers"))
}

func runError() {
	fmt.Println("clog.Error():")
	run(false)
}

func runSuccess() {
	fmt.Println("clog.Ok():")
	run(true)
}

func run(success bool) {
	data, err := getData(success)
	if err != nil {
		Fatallb(err, "calling getData(false) from Test()")
		return
	}
	Ok(data, "Data received")
}

func getData(success bool) ([]interface{}, error) {
	if !success {
		return nil, errors.New("something very specific was wrong")
	}
	return []interface{}{6, "Charizard"}, nil
}
