# wc

## Overview

In this project, I build my own version of the Unix command line tool wc in Golang!

The functional requirements for wc are concisely described by its man page - give it a go in your local terminal now:

```text
 man wc
```

The TL/DR version is: wc – word, line, character, and byte count

## Getting Started
Follow the steps below to get started with the wc program:
    
1. Clone the repository using Git:
   ```text
   git clone https://github.com/ktran1005/wc
   ```
2. Change to the project directory:
   ```text
   cd wc
   ```
3. Build the project:
   ```text
   go build wc.go
   ```

## Usages

```text
./wc [OPTION] [FILE]

OPTION:
    -c, --bytes
        print the byte counts
    -m, --chars
        print the character counts. If the current locale does not support
        multibyte characters, this will match the -c option
    -l, --lines
        print the newline counts
    -w, --words
        print the word counts
```
The program also supports to read from standard input if no filename is specified
 ```text
 cat [FILE] | ./wc [OPTION]
 ```


## Default
By default, the program outputs the number of words, lines, bytes, and characters for the file.

## Demo
![image](https://github.com/ktran1005/wc/assets/88155108/ae4c3252-25f3-4773-b67d-d24ef13b950d)

   

