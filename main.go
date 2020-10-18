package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var repeat = 10
var num = 1
var totalTime = 100000

func main() {
	rand.Seed(time.Now().Unix())
	sum := 0.0

	for i := 0; i < repeat; i++ {
		itv := make([]int, 0)
		for j := 0; j < num; j++ {
			itv = append(itv, rand.Intn(15)+286)
		}
		sort.Ints(itv)
		time := totalTime
		butterUntil := -1
		curr := itv[0]
		for curr <= time {
			if rand.Float64() <= 0.25 {
				if butterUntil == -1 {
					time += 400
				} else {
					time += curr - (butterUntil - 400)
				}
				butterUntil = curr + 400
			} else {
				butterUntil = -1
			}

			if len(itv) == 1 {
				itv = nil
			} else {
				itv = itv[1:]
			}
			//fmt.Println(itv)
			itv = append(itv, curr+rand.Intn(15)+286)
			sort.Ints(itv)
			//fmt.Println(itv)
			curr = itv[0]
		}
		sum += float64(time) / 100.0
	}

	fmt.Println("Avg = ", (sum*100/float64(repeat))/float64(totalTime))

}
