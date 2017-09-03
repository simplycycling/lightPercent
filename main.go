package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	a, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/brightness")
	check(err)
	fmt.Println(string(a))
}
