package main

import (
	"fmt"
	"strconv"
	"time"
)

type person struct {
	name  string
	value int64
}

func main() {
	clock := person{value: 0, name: "diogoxiang"}
	fmt.Printf("%v", &clock)
	// fmt.Printf("%d\n", clock.epoch)
	// fmt.Printf("%d\n", time.Now().Unix())

	// t := time.Now().Unix()
	// str := strconv.FormatInt(t, 10)
	// fmt.Println(str)

	// var _startDate int64 = time.Now().Unix()
	// // var startDate string = time.Unix(_startDate, 0).Format("2006-01-02 15:04:05")
	// fmt.Println(_startDate)

	_startDate := strconv.FormatInt(1501214351, 10)
	fmt.Println(_startDate)
}

type clock interface {
	epoch() int64
}

type systemClock struct {
}

func (c systemClock) epoch() int64 {
	return time.Now().Unix()
}
