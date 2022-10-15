package main

import "fmt"

func kmpAlgorithm(str string, pattern string) bool {
	LPSArr := buildLPSArr(pattern)
	return stringMatching(str, pattern, LPSArr)
}

func buildLPSArr(pattern string) []int {
	LPSArr := []int{}
	for i := 0; i < len(pattern); i++ {
		LPSArr = append(LPSArr, -1)
	}
	fmt.Println("Init LPS array", LPSArr)

	i, j := 1, 0
	for i < len(pattern) {
		fmt.Printf("j=%d, i=%d, LPS array=%+v\n", j, i, LPSArr)
		if pattern[i] == pattern[j] {
			LPSArr[i] = j
			i += 1
			j += 1
		} else if j > 0 {
			j = LPSArr[j-1] + 1
		} else {
			i += 1
		}
	}
	return LPSArr
}

func stringMatching(str string, pattern string, LPSArr []int) bool {
	i, j := 0, 0
	for i+len(pattern)-j <= len(str) {
		if str[i] == pattern[j] {
			if j == len(pattern)-1 {
				return true
			}
			i += 1
			j += 1
		} else if j > 0 {
			j = LPSArr[j-1] + 1
		} else {
			i += 1
		}
	}
	return false
}

func main() {
	str := "kmppkmpkmpkmdkmpkmpk"
	pattern := "kmpkmdkmpkmpk"
	fmt.Println("does match:", kmpAlgorithm(str, pattern))
}

/*
output:
Init LPS array [-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1]
j=0, i=1, LPS array=[-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1]
j=0, i=2, LPS array=[-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1]
j=0, i=3, LPS array=[-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1]
j=1, i=4, LPS array=[-1 -1 -1 0 -1 -1 -1 -1 -1 -1 -1 -1 -1]
j=2, i=5, LPS array=[-1 -1 -1 0 1 -1 -1 -1 -1 -1 -1 -1 -1]
j=0, i=5, LPS array=[-1 -1 -1 0 1 -1 -1 -1 -1 -1 -1 -1 -1]
j=0, i=6, LPS array=[-1 -1 -1 0 1 -1 -1 -1 -1 -1 -1 -1 -1]
j=1, i=7, LPS array=[-1 -1 -1 0 1 -1 0 -1 -1 -1 -1 -1 -1]
j=2, i=8, LPS array=[-1 -1 -1 0 1 -1 0 1 -1 -1 -1 -1 -1]
j=3, i=9, LPS array=[-1 -1 -1 0 1 -1 0 1 2 -1 -1 -1 -1]
j=4, i=10, LPS array=[-1 -1 -1 0 1 -1 0 1 2 3 -1 -1 -1]
j=5, i=11, LPS array=[-1 -1 -1 0 1 -1 0 1 2 3 4 -1 -1]
j=2, i=11, LPS array=[-1 -1 -1 0 1 -1 0 1 2 3 4 -1 -1]
j=3, i=12, LPS array=[-1 -1 -1 0 1 -1 0 1 2 3 4 2 -1]
does match: true
*/
