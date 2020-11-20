package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 外层循环次数（防假死）
var outerRepeat = 10

// 内层循环次数
var repeat = 100000

// 玉米个数
var num = 5

// 基准时间
var basicTime = 100000

func main() {
	avg := 0.0
	for k := 0; k < outerRepeat; k++ {
		rand.Seed(time.Now().Unix())
		sum := 0.0
		totalTime := basicTime - 1000 + rand.Intn(1000)

		for i := 0; i < repeat; i++ {

			// 录入第一次玉米攻击时机
			itv := make([]int, 0)
			for j := 0; j < num; j++ {
				itv = append(itv, rand.Intn(15)+286)
			}
			sort.Ints(itv)

			time := totalTime
			butterUntil := -1
			curr := itv[0]
			for curr <= time {
				if curr >= butterUntil {
					butterUntil = -1
				}
				if rand.Float64() <= 0.25 {
					if butterUntil == -1 {
						time += 400
					} else {
						time += curr - (butterUntil - 400)
					}
					butterUntil = curr + 400
				}

				if len(itv) == 1 {
					itv = nil
				} else {
					itv = itv[1:]
				}
				itv = append(itv, curr+rand.Intn(15)+286)
				sort.Ints(itv)
				curr = itv[0]
			}
			sum += float64(time) / 100.0
		}
		result := (sum * 100 / float64(repeat)) / float64(totalTime)
		fmt.Println(result)
		avg += result
	}
	fmt.Println("Repeat = ", outerRepeat*repeat, " times")
	fmt.Println("Avg = ", avg/float64(outerRepeat))
}
