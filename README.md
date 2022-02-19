# bible
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/bible)](https://goreportcard.com/badge/github.com/Jmainguy/bible)
[![Release](https://img.shields.io/github/release/Jmainguy/bible.svg?style=flat-square)](https://github.com/Jmainguy/bible/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/bible/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/bible?branch=main)

A command line bible written in go

## Usage
```/bin/bash
  -compare string
    	Translations to compare passage against, set to 'all' for all translations in the database, or use a space seperated list
  -db string
    	Bible database to use (defaults to ~/.bible/bible.db, will download this for you if it doesnt exist)
  -generateTests
    	Whether to generate and print tests to stdout
  -listBooks
    	List all books of the bible and their number of chapters
  -listTranslations
    	List all translations in database
  -passage string
    	Passage to return. Can be given in following syntax. 'John', '1 John 3', 'John 3:16', or for a range in the same book '1 John 1:1 - 3:16' (default "John 3:16")
  -randomChapter
    	Print a random Chapter
  -randomPassage
    	Print a random passage
  -translation string
    	Bible translation to use (default "t_kjv")
```

## Example
Single verse from the NASB

```/bin/bash
[jmainguy@jmainguy bible]$ bible  --passage "Obadiah 1:21" -translation t_nasb
Obadiah 1:21 - New American Standard Bible
  The deliverers will ascend Mount Zion To judge the mountain of Esau, And the kingdom will be the Lord’s.

```

Compare multiple verses, against multiple translations.

```/bin/bash
[jmainguy@jmainguy bible]$ bible  --passage "Philippians 3:13 - 3:14" --compare "t_kjv t_leb t_nasb"
Philippians 3:13 - King James Version
  Brethren, I count not myself to have apprehended: but this one thing I do, forgetting those things which are behind, and reaching forth unto those things which are before,

Philippians 3:14 - King James Version
  I press toward the mark for the prize of the high calling of God in Christ Jesus.

Philippians 3:13 - Lexham English Bible
  Brothers, I do not consider myself to have laid hold of it. But I do one thing, forgetting the things behind and straining toward the things ahead,

Philippians 3:14 - Lexham English Bible
  I press on toward the goal for the prize of the upward call of  God in Christ Jesus.

Philippians 3:13 - New American Standard Bible
  Brethren, I do not regard myself as having laid hold of it yet; but one thing I do: forgetting what lies behind and reaching forward to what lies ahead,

Philippians 3:14 - New American Standard Bible
  I press on toward the goal for the prize of the upward call of God in Christ Jesus.
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/bible/releases)

## Install

### Homebrew

```/bin/bash
brew install jmainguy/tap/bible
```

### yum / dnf
You can install the rpm from the [The Releases Page](https://github.com/Jmainguy/bible/releases)
```/bin/bash
# For example, to install 0.0.0-1
dnf install https://github.com/Jmainguy/bible/releases/download/v0.0.0/bible-0.0.0-1.x86_64.rpm
```

### deb
There is a debian package in [The Releases Page](https://github.com/Jmainguy/bible/releases) 

```
# For example, to install 0.0.1
wget https://github.com/Jmainguy/bible/releases/download/v0.0.1/bible_x86_64.deb
dpkg -i bible_x86_64.deb
```

## Build
```/bin/bash
export GO111MODULE=on
go build
```

