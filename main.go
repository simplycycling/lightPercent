package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func percentage(current int) float64 {
	const max = 937
	percent := (float64(current) * float64(100)) / float64(max)
	return percent
}

func main() {
	v, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/brightness")
	check(err)
	s := string(v)
	i, err := strconv.Atoi(s)
	check(err)
	fmt.Println(percentage(i))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
