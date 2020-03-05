package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"

	"github.com/spf13/afero"
)

var startWhitespace = regexp.MustCompile(`^\s+(?:\S)`)

func main() {
	fs := afero.NewOsFs()
	file, err := fs.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		line = fixWhitespace(line)
		fmt.Println(string(line))
	}
}

func fixWhitespace(line []byte) []byte {
	numSpaces := numLeadingWhitespaces(line)
	return whitespaceFmt(numSpaces, line)
}

func numLeadingWhitespaces(line []byte) int {
	var numWhitespace int
	result := startWhitespace.FindSubmatch(line)
	if len(result) != 0 {
		fullMatch := result[0]
		numWhitespace = len(fullMatch) - 1
	}
	return numWhitespace
}

func whitespaceFmt(numLeadingSpaces int, line []byte) []byte {
	withoutSpaces := line[numLeadingSpaces:]

	freshLine := make([]byte, len(withoutSpaces)+numLeadingSpaces)

	for i := 0; i < roundToNearest4(numLeadingSpaces); i++ {
		freshLine[i] = ' '
	}
	return append(freshLine, withoutSpaces...)
}

func roundToNearest4(in int) int {
	return 4 * int(math.Round(float64(in)/4))
}
