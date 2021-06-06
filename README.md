# bible
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/bible)](https://goreportcard.com/badge/github.com/Jmainguy/bible)
[![Release](https://img.shields.io/github/release/Jmainguy/bible.svg?style=flat-square)](https://github.com/Jmainguy/bible/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/bible/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/bible?branch=main)

A command line bible written in go

## Usage
```/bin/bash
Usage of ./bible:
  -compare string
    	Translations to compare passage against, set to 'all' for all translations in the database, or use a space seperated list
  -db string
    	Bible database to use (default "database/bible-sqlite-jmainguy.db")
  -generateTests
    	Whether to generate and print tests to stdout
  -listBooks
    	List all books of the bible and their number of chapters
  -listTranslations
    	List all translations in database
  -passage string
    	Passage to return. Can be given in following syntax. 'John', '1 John 3', 'John 3:16', or for a range in the same book '1 John 1:1 - 3:16' (default "John 3:16")
  -translation string
    	Bible translation to use (default "t_kjv")
```

## Example
```/bin/bash
[jmainguy@jmainguy bible]$ ./bible  --passage "Obadiah 1:21" -translation t_nasb
Obadiah 1:21 - New American Standard Bible
  The deliverers will ascend Mount Zion To judge the mountain of Esau, And the kingdom will be the Lord’s.

```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/bible/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```

