package main

import (
	"flag"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

func setUpTest() {
	os.Args = append(os.Args, "bible")
}

func tearDownTest() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = nil
}

func TestRunMain(t *testing.T) {
	setUpTest()

	main()
	tearDownTest()
}

func TestRunMainRandomPassage(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--randomPassage")
	main()
	tearDownTest()
}

func TestRunMainRandomChapter(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--randomChapter")
	main()
	tearDownTest()
}

func TestRunMainListBooks(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--listBooks")
	main()
	tearDownTest()
}

func TestRunMainListTranslations(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--listTranslations")
	main()
	tearDownTest()
}

func TestRunMainEsther(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--passage")
	os.Args = append(os.Args, "Esther")
	main()
	tearDownTest()
}

func TestRunMainEstherLeb(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--passage")
	os.Args = append(os.Args, "Esther")
	os.Args = append(os.Args, "-translation")
	os.Args = append(os.Args, "t_leb")
	main()
	tearDownTest()
}

func TestRunMainCompareEstherLebNasb(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--passage")
	os.Args = append(os.Args, "Esther")
	os.Args = append(os.Args, "-translation")
	os.Args = append(os.Args, "t_leb")
	os.Args = append(os.Args, "--compare")
	os.Args = append(os.Args, "t_leb")
	os.Args = append(os.Args, "t_nasb")

	main()
	tearDownTest()
}

func TestRunMainCompareEstherAll(t *testing.T) {
	setUpTest()

	os.Args = append(os.Args, "--passage")
	os.Args = append(os.Args, "Esther")
	os.Args = append(os.Args, "-translation")
	os.Args = append(os.Args, "t_leb")
	os.Args = append(os.Args, "--compare")
	os.Args = append(os.Args, "all")

	main()
	tearDownTest()
}
