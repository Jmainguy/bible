package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadDB(file string) error {

	// Ensure dir exists, if not create
	dirpath := filepath.Dir(file)
	dir, err := os.Stat(dirpath)
	if err != nil {
		err = os.MkdirAll(dirpath, os.ModePerm)
		if err != nil {
			fmt.Printf("So, %s did not exist, we tried to download and create it, but ran into an error\n", file)
			fmt.Printf("Specifically, %s the directory did not exist, and we got this error when trying to create it\n", dirpath)
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if !dir.IsDir() {
			err = os.Remove(dirpath)
			if err != nil {
				fmt.Printf("So, %s did not exist, we tried to download and create it, but ran into an error\n", file)
				fmt.Printf("Specifically, %s existed but was not a directory, so we tried to delete it, then recreate it as a dir but ran into this error\n", dirpath)
				fmt.Println(err)
				os.Exit(1)
			}
			err = os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				fmt.Printf("So, %s did not exist, we tried to download and create it, but ran into an error\n", file)
				fmt.Printf("Specifically, %s existed but was not a directory, so we tried to delete it, then recreate it as a dir but ran into this error\n", dirpath)
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	// Get the data
	resp, err := http.Get("https://github.com/Jmainguy/bible/raw/main/database/bible.db")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}
