package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/dpritchett/gubs"
)

func trReader(reader *bufio.Reader) {

	for {
		line, err := reader.ReadString('\n')
		if err == nil || err == io.EOF {
			fmt.Print(gubs.Tr(line))
		}

		if err != nil {
			break
		}
	}
}

func promptOnceIfInteractive() {
	if termutil.Isatty(os.Stdin.Fd()) {
		fmt.Print("Ⓛ ⓔ ⓣ 'ⓢ  ⓜ ⓐ ⓚ ⓔ  ⓢ ⓞ ⓜ ⓔ  ⓑ ⓤ ⓑ ⓑ ⓛ ⓔ ⓢ !\n> ")
	}
}

// StripStdin reads text from SDTIN and emits that same text minus any
// ANSI escape codes.
func TrStdin() {
	promptOnceIfInteractive()
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
