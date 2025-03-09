package main

import (
	"cut/processor"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 5 {
		fmt.Println("Usage: ./cut -f<field> [-d<delimiter>] <file>")
		return
	}

	var delimiter string
	var fileName string
	var fieldNum int
	fieldNum = -1

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "-f") {
			num, err := strconv.Atoi(arg[2:])
			if err != nil {
				fmt.Println("Invalid field number")
				return
			}
			fieldNum = num
		} else if strings.HasPrefix(arg, "-d") {
			parts := strings.SplitN(arg, "-d", 2)
			if len(parts) < 2 || parts[1] == "" {
				fmt.Println("Invalid delimiter")
				return
			}
			delimiter = parts[1]
		} else {
			fileName = arg
		}
	}

	if fieldNum == -1 {
		fmt.Println("Field number (-f<field>) is required")
		return
	}

	if delimiter == "" {
		delimiter = "\t"
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	processor.Processor(file, delimiter, fieldNum)
}
