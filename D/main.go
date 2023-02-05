package main

import (
	"fmt"
	"strconv"
	"strings"
)

var ans map[int]bool

func diap(s, e int) {
	for i := s; i <= e; i++ {
		ans[i] = true
	}
}

func main() {
	var myCount int
	fmt.Scan(&myCount)
	for i := 0; i < myCount; i++ {
		var count int
		fmt.Scan(&count)
		var str string
		fmt.Scan(&str)
		s := make(map[int]bool)
		ans = s
		for i := 1; i <= count; i++ {
			ans[i] = false
		}
		var result []string
		pages := strings.Split(str, ",")
		for _, page := range pages {
			if len(strings.Split(page, "-")) == 2 {
				buf := strings.Split(page, "-")
				buf1, _ := strconv.Atoi(buf[0])
				buf2, _ := strconv.Atoi(buf[1])
				diap(buf1, buf2)
			} else {
				buf, _ := strconv.Atoi(page)
				ans[buf] = true
			}
		}
		for i := 1; i <= count; i++ {
			if ans[i] == true {
				continue
			}
			if ((ans[i+1] == true) && i+1 <= count) || i == count {
				result = append(result, fmt.Sprint(i))
			} else {
				var end int
				for j := i; j <= count; j++ {
					end = j
					if ans[j+1] == true && j+1 <= count {
						break
					}
				}
				result = append(result, fmt.Sprintf("%d-%d", i, end))
				i = end + 1
			}
		}
		for i, s2 := range result {
			if len(result) == i+1 {
				fmt.Print(s2)
				break
			}
			fmt.Printf("%s,", s2)
		}
		fmt.Printf("\n")

	}
}
