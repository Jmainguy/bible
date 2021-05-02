# bible
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/bible)](https://goreportcard.com/badge/github.com/Jmainguy/bible)
[![Release](https://img.shields.io/github/release/Jmainguy/bible.svg?style=flat-square)](https://github.com/Jmainguy/bible/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/bible/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/bible?branch=main)

A command line bible written in go

## Usage
```/bin/bash
Usage of ./bible:
  -db string
    	Bible database to use (default "database/bible-sqlite-jmainguy.db")
  -verse string
    	Verse to return. Can be given in following syntax. 'John', '1 John 3', 'John 3:16', or for a range in the same book '1 John 1:1 - 3:16' (default "John 3:16")
  -version string
    	Bible Version to use (default "t_kjv")
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/bible/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```

