package main

import "fmt"

func stringMatching(str string, pattern string) bool {
	for i := 0; i < len(str); i++ {
		j := 0
		for j < len(pattern) {
			if str[i+j] != pattern[j] {
				break
			}
			j += 1
		}
		if j == len(pattern) {
			return true
		}
	}
	return false
}

func main() {
	str := "kmppkmpkmpkmdkmpkmpk"
	pattern := "kmpkmdkmpkmpk"
	fmt.Println("does match:", stringMatching(str, pattern))
}

/*
output:
does match: true
*/
