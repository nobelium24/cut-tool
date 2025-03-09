package processor

import (
	"bufio"
	"fmt"
	"os"
)

func Processor(file *os.File, delimiter string, fieldNum int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
