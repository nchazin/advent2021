package main

import "fmt"

func makefish(fishschool map[int]int, days int) map[int]int {
	fmt.Printf("START Day: %d len %v\n", 0, fishschool)
	for day := 1; day <= days; day++ {
		newschool := make(map[int]int)
		for timer, count := range fishschool {
			if timer == 0 {
				//fmt.Println("Hello ", count)
				newschool[6] += count
				newschool[8] += count
			} else {
				newschool[timer-1] += count
			}

		}
		fishschool = newschool
		fmt.Printf("Day: %d len %v\n", day, fishschool)
	}

	return fishschool
}

func main() {
	fishschool := make(map[int]int)
	fishschool[3] = 2
	fishschool[4] = 1
	fishschool[1] = 1
	fishschool[2] = 1

	finalschool := makefish(fishschool, 18)
	fmt.Println(finalschool)
	fmt.Println(count_fish(finalschool))

	finalschool = makefish(fishschool, 80)
	fmt.Println(finalschool)
	fmt.Println(count_fish(finalschool))
}

func count_fish(fishschool map[int]int) int {
	size := 0

	for _, count := range fishschool {
		size += count
	}
	return size
}
