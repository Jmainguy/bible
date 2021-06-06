package main

import (
	"fmt"
	"os"
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

func dbCheck(e error) {
	if e != nil {
		if strings.Contains(e.Error(), "no such table:") {
			fmt.Println(e)
			fmt.Println("This usually means your database doesnt exist, or is corrupt")
			os.Exit(1)
		}
	}
}
