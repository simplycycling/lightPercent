package main

import (
	"fmt"

)

//maxBrightness := 937
//current := func now() int {
//	a, err := ioutil.ReadFile("/sys/class/backlight/intel_backlight/brightness")
//	check(err)
//	s := string(a)
//	i, err := strconv.Atoi(s)
//	return i
//}

func percentage(current int) float64 {
	const max = 937
	percent := (float64(current) * float64(100)) / float64(max)
	return percent
}

func main() {
	fmt.Println(percentage(504))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
