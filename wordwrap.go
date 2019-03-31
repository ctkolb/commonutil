package commonutil

import (
	"bytes"
	"strings"
	"unicode"
)

func WordWrap(s string, limit int) string {
	//changes whitespace in a string into \n as needed to wrap the output such that it
	//	doesn't exceed a specified limit (width), while not breaking the text in the middle
	//	of a word

	//if the input string is smaller than the limit, we early exit
	if len(s) <= limit {
		return s
	}

	//split into words on whitespace using fields
	words := strings.Fields(s)
	//convert the string into a rune array so we can grab certain runes as needed
	runes := []rune(s)
	//we'll use two buffers: currentLine is just the line we're working with, and messageBuffer is the one we dump the full message into
	var currentLine, messageBuffer bytes.Buffer
	//we count every character we process so that we can find what comes next in the runes array
	var index int

	//start by adding any whitespace that occurs at the beginning
	for _, r := range runes {
		if unicode.IsSpace(r) {
			messageBuffer.WriteRune(r)
			index = index + 1
		} else {
			// if we hit something that isn't whitespace we are done here
			break
		}
	}
	//range through the words
	for _, w := range words {
		//if there's anything already on the line, and the currentLine buffer + the word length is less than the limit width
		if currentLine.Len() > 0 && currentLine.Len()+len(w) >= limit {
			//	Add the currentLine buffer + \n to the messageBuffer
			messageBuffer.WriteString(currentLine.String() + "\n")
			//	clear currentLine
			currentLine.Reset()
		}
		//Add the word into the current line buffer
		currentLine.WriteString(w)
		index = index + len(w)

		//now we deal with white space that happens after that last word
		//we have to also deal with the fact that there may be more than one whitespace rune
		for index <= len(runes)-1 && unicode.IsSpace(runes[index]) {
			//grab the whitespace rune that occurs after this word and hold on to it
			whiteSpace := runes[index]
			//Find the whitespace that corresponds to what come right after this word and add it
			currentLine.WriteRune(whiteSpace)
			index = index + 1
			//if the whitespace is a newline, then we need to reset the currentLine buffer after this to reset the count
			if whiteSpace == '\n' {
				//dump the current line into the messagebuffer
				messageBuffer.WriteString(currentLine.String())
				// clear currentLine
				currentLine.Reset()
			}
		}
	}
	//after the last word, we still need to dump currentLine in
	messageBuffer.WriteString(currentLine.String())
	//add any whitespace that is at the end of the runes
	for i := index; i < len(runes); i++ {
		messageBuffer.WriteRune(runes[i])
	}

	return messageBuffer.String()
}
