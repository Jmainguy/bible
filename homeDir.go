package main

import (
	"fmt"
	"os"
)

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	return dirname
}
