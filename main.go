package main

import (
	"alien-numeral/alien"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter alien numeral: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := alien.AlienToInt(input)
	fmt.Printf("Converted to integer: %d\n", result)
}
