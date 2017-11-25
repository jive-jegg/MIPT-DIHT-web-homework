package main

import "strings"

func RemoveEven(a []int) []int {
	s := make([]int, 0, len(a))
	for _, val := range a {
		if val % 2 == 1 {
			s = append(s, val)
		}
	}
	return s
}


func PowerGenerator(num int) func()int {
	ans := 1
	return func() int {
		ans = ans * num
		return ans
	}
}


func DifferentWordsCount(s string) int {
	s = strings.ToLower(s)
	diffWords := make([]string, 0, len(s))
	lastIdx := 0
	for i := 0; i < len(s); i++ {
		if !('a' <= s[i] && s[i] <= 'z') {
			if lastIdx != i {
				j := 0
				newSlice := s[lastIdx:i]
				for j < len(diffWords) && diffWords[j] != newSlice {
					j++
				}
				if j == len(diffWords) {
					diffWords = append(diffWords, newSlice)
				}
			}
			lastIdx = i + 1
		}
	}
	if lastIdx != len(s) {
		j := 0
		newSlice := s[lastIdx:len(s)]
		for j < len(diffWords) && diffWords[j] != newSlice {
			j++
		}
		if j == len(diffWords) {
			diffWords = append(diffWords, newSlice)
		}
	}
	return len(diffWords)
}