package api

import (
	"strconv"
)

func FizzBuzz(int1, int2, limit int, str1, str2 string) []string {
	res := make([]string, 0)
	for i := 1; i <= limit; i++ {
		var iteration string
		if int1*int2 != 0 && i%(int1*int2) == 0 {
			iteration = str1 + str2
		} else if int1 != 0 && i%int1 == 0 {
			iteration = str1
		} else if int2 != 0 && i%int2 == 0 {
			iteration = str2
		} else {
			iteration = strconv.Itoa(i)
		}
		res = append(res, iteration)
	}
	return res
}
