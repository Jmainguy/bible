package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestSetup(t *testing.T) {
	setUpTest()

	_, err := setup()
	assert.Equal(t, err, nil)
	tearDownTest()
}

func TestSetupNoHome(t *testing.T) {
	home := os.Getenv("HOME")
	os.Setenv("HOME", "")

	setUpTest()

	_, err := setup()
	assert.Equal(t, err, fmt.Errorf("$HOME is not defined"))
	tearDownTest()
	os.Setenv("HOME", home)

}

func TestSetupNoDB(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--db")
	os.Args = append(os.Args, "/tmp/bible.db")

	_, err := setup()
	assert.Equal(t, err, nil)
	tearDownTest()
	os.Remove("/tmp/bible.db")
}

func TestSetupBadDB(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--db")
	os.Args = append(os.Args, "/proc/bible.db")

	_, err := setup()

	assert.NotEqual(t, err, nil)
	tearDownTest()

}

func TestSetupDBDevNull(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--db")
	os.Args = append(os.Args, "/dev/null")

	_, err := setup()

	assert.NotEqual(t, err, nil)
	tearDownTest()

}
