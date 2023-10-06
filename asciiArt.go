package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// w := http.ResponseWriter (lets code Fprint to html); banner := font text file, as chosen by radio buttons; textbox := input in textbox on html
func asciiArt(w http.ResponseWriter, banner string, textbox string, output *os.File) {

	// standard.txt (readData), split by '\n'
	splitStr := strings.Split(string(readData(banner+".txt")), "\n")

	// replaces new lines (\n) in textbox with '\\n'
	replaceNewline := strings.ReplaceAll(textbox, "\r\n", "\\n")
	// Textbox, split by '\\n', literally the symbols \ and n together
	splitText := strings.Split(string(replaceNewline), "\\n")
	// for each slice of arg...
	for a := 0; a < len(splitText); a++ {

		runeArgs := []rune(splitText[a])
		// if the slice of arg contains nothing then print a new line

		for i := 1; i <= 8; i++ {
			for j := 0; j <= len(runeArgs)-1; j++ {
				letterArgs := runeArgs[j]
				// rune to line number
				lineNumber := (int(letterArgs)-32)*9 + i
				// prints from line number to the next 8 lines
				//	fmt.Println(splitStr[lineNumber])
				//	fmt.Fprint(w, splitStr[lineNumber])

				fmt.Fprint(output, splitStr[lineNumber])

			}
			//fmt.Fprintln(w)
			fmt.Fprintln(output)
		}
	}
	// fmt.Fprintln( "<button type="submit" formaction="/download">
	// download</button>")
}
