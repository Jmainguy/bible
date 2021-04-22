package main

import (
	"fmt"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func checkNumeric(s string) bool {
	alphabet := "0123456789"
	for _, character := range s {
		if !strings.Contains(alphabet, string(character)) {
			return false
		}
	}
	return true
}

func checkString(s string) bool {
	alphabet := "0123456789:-"
	for _, character := range s {
		if !strings.Contains(alphabet, string(character)) {
			return false
		}
	}
	return true
}
