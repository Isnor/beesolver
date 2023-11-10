# Beesolver

Provides solutions to the New York Times "spelling bee" game.

## Overview

This is just a silly project I hacked together one night when I was too stupid to figure out the word properly. It's very simple and doesn't have a ton of thought put into it.

## Usage

Build the binary with `make` and run `./beesolver` to view the help:
```
NAME:
   beesolver

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -l value        The 6 letters of the puzzle
   --middle value  the middle, required letter
   --dict value    path to the words list (default: "/usr/share/dict/words")
   --help, -h      show help
```

Example:

`./bsolve --middle o --letters d,m,t,r,y,i`

This probably only works on Linux, because by default it tries to read a "dictionary" of words from `/usr/share/dict/words`. It might work on a Mac, but otherwise, an alternative word list can be provided via the `--dict` argument. There is a mega word list provided in this repository that I got from `https://github.com/dwyl/english-words` that can be used for this purpose:

`./bsolve --middle o --dict ./words_alpha.txt --letters d,m,t,r,y,i`

