package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadDB(file string) (retErr error) {

	// Ensure dir exists, if not create
	dirpath := filepath.Dir(file)
	dir, err := os.Stat(dirpath)
	if err != nil {
		err = os.MkdirAll(dirpath, os.ModePerm)
		if err != nil {
			retErr = fmt.Errorf("tried to create dir %s, got error %s", dirpath, err)
			return retErr
		}
	} else {
		if !dir.IsDir() {
			err = os.Remove(dirpath)
			if err != nil {
				retErr = fmt.Errorf("tried to delete %s, got error %s", dirpath, err)
				return retErr
			}
			err = os.MkdirAll(dirpath, os.ModePerm)
			if err != nil {
				retErr = fmt.Errorf("tried to create dir %s, got error %s", dirpath, err)
				return retErr
			}
		}
	}

	// Get the data
	resp, err := http.Get("https://github.com/Jmainguy/bible/raw/main/database/bible.db")
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("error closing response body: %v\n", err)
		}
	}()

	// Create the file
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			fmt.Printf("error closing output file: %v\n", err)
		}
	}()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}
