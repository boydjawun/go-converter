package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("   converter <decimal>")
		fmt.Println("   converter -x <hex>")
		fmt.Println("   converter -b <binary>")
		os.Exit(1)
	}

	var base int
	var numStr string

	switch os.Args[1] {
	case "-x":
		base = 16
		numStr = os.Args[2]
	case "-b":
		base = 2
		numStr = os.Args[2]
	default:
		base = 10
		numStr = os.Args[1]
	}

	n, err := strconv.ParseInt(numStr, base, 64)
	if err != nil {
		fmt.Println("Error: invalid number: ", numStr)
	}

	fmt.Printf("Decimal: %d\n", n)
	fmt.Printf("Hex:     0x%X\n", n)
	fmt.Printf("Binary:  0b%b\n", n)
}

// Turn to an executable
/**
1. go build -o converter.exe converter.go
2. create a folder to hold the executable
3. move the exe into the folder: move <executable>.exe <executable-folder>
4. Add folder to PATH list in environment variables
5. Restart terminal
*/
