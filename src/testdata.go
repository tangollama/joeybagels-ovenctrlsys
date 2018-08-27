package main

import (
	"math/rand"
	"time"
	"fmt"
)
/*
Simulate() returns a JSON object in the following form

{
	"temp": 450,		//temperature of the oven
	"version": "1.0.4",	//software version of the control software
	"restart": false 	//boolean of restart Event
}
*/

type Status struct {
	Temp int `json:"temp"`
	Version string `json:"version"`
	Restart bool `json:"restart"`
}

func Simulate() (output string) {

	temps := []int {450,500,400,450,375,425,375,400,450,525,650,325,350,375,400}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(temps)-1)
	restart := index%8 == 0	//mark restart on a random index of 8
	output = fmt.Sprintf(`{ "temp": %d, "version": "%s", "restart": %t }`, temps[index], "1.0.1", restart)
	return output
}

/*func main() {
	v := Simulate()
	fmt.Println(v)
}*/

