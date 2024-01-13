package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"log"
	"unicode"
	"strings"
	"errors"
)

func checkFileExists(filePath string) bool {
    _, error := os.Stat(filePath)

    return !errors.Is(error, os.ErrNotExist)
}

func handleParams() (string, string, error) {
    var fileName string
    var option string
    if len(os.Args) == 3 {
        option = string(os.Args[1])
        if  !(option == "-c" || option == "-l" || option == "-w" || option == "-m") {
            fmt.Println("Unknown option -c. Please pass one of these options (-c/-l/-w/-m)")
            os.Exit(-1)
        }
        
        isFileExist := checkFileExists(string(os.Args[2]))
        if isFileExist {
            fileName = os.Args[2]
            option = os.Args[1]
        } else {
            fmt.Printf("File does not exist. Please provide the correct input file\n")
            os.Exit(-1)
        }        
    } else if len(os.Args) == 2 {
        isFileExist := checkFileExists(string(os.Args[1]))
        if isFileExist {
            fileName = os.Args[1]
            option = ""
        } else {
            fmt.Printf("File does not exist. Please provide the correct input file\n")
            os.Exit(-1)
        }
    } else {
        fmt.Printf("Usage: ./ccwc -option[-c/-l/-w/-m] input_file")
        os.Exit(-1)
    }
    return fileName, option, nil
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkLocale() bool {
    locale := os.Getenv("LANG")
    if strings.Contains(locale, "UTF-8") {
        return true
    }
    return false
}

func getFileInfo(file io.Reader) (int, int, int, int) {
	bytes := 0
	characters := 0
	lines := 0
	words := 0
	inWord := false
	locale := checkLocale()
	r := bufio.NewReader(file)
	for {
		if c, sz, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				if inWord {
					words++
				}
				bytes += int(sz)
				break
			}
		
		} else {
			if c == '\n' {
				lines++
			}
			if unicode.IsSpace(c) {
				if inWord {
					words++
				}
				inWord = false
			} else {
				inWord = true
			}

			bytes += sz
			if locale == true {
				characters++
			} else {
				characters += sz
			}
		}
	}
	return bytes, characters, lines, words
} 

func main() {
	var option string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, characters, lines, words := getFileInfo(os.Stdin)		
	
		if len(os.Args) < 2 {
			option = ""
		} else {
			option = os.Args[1]
		}

		switch option {
		case "-c":
			fmt.Printf("%v\n", bytes)
		case "-m": 
			fmt.Printf("%v\n", characters)		
		case "-l": 
			fmt.Printf("%v\n", lines)
		case "-w": 
			fmt.Printf("%v\n", words)
		default:
			fmt.Printf("   %v   %v  %v\n", lines, words, bytes)
		}

	} else {
		fileName, option, err := handleParams()
		check(err)
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		bytes, characters, lines, words := getFileInfo(file)

		switch option {
		case "-c":
			fmt.Printf("%v %v\n", bytes, fileName)
		case "-m": 
			fmt.Printf("%v %v\n", characters, fileName)		
		case "-l": 
			fmt.Printf("%v %v\n", lines, fileName)
		case "-w": 
			fmt.Printf("%v %v\n", words, fileName)
		default:
			fmt.Printf("%v %v %v %v\n", lines, words, bytes, fileName)
		}
	}
}