package main

import (
	"fmt"
	"os"
	"strings"
)

func checkNumeric(s string) bool {
	alphabet := "0123456789"
	for _, character := range s {
		if !strings.Contains(alphabet, string(character)) {
			return false
		}
	}
	return true
}

func checkNumericPlusColonDash(s string) bool {
	alphabet := "0123456789:-"
	for _, character := range s {
		if !strings.Contains(alphabet, string(character)) {
			return false
		}
	}
	return true
}

func checkErrorPrintExitOne(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
