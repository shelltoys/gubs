package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dpritchett/gubs"
)

func trReader(reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			print(gubs.Tr(line) + "\n")
		}

		if err != nil {
			break
		}
	}
}

// StripStdin reads text from SDTIN and emits that same text minus any
// ANSI escape codes.
func TrStdin() {
	reader := bufio.NewReader(os.Stdin)
	trReader(reader)
}

func main() {
	if len(os.Args) > 1 {
		for _, str := range os.Args[1:] {

			fmt.Print(gubs.Tr(str) + " ")
		}
	} else {
		TrStdin()
	}
}
