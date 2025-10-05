package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	buffer := make([]byte, 8) // read 8 bytes at a time
	currentLine := ""        // holds the incomplete line between reads

	for {
		n, err := f.Read(buffer)
		if err == io.EOF {
			// end of file → break out after loop
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}

		chunks := string(buffer[:n])
		parts := strings.Split(chunks,"\n")

		for i :=0 ; i < len(parts)-1 ; i++ {
			fmt.Printf("read: %s \n",currentLine+parts[i])
			currentLine = ""
		}

		currentLine+=parts[len(parts)-1]
	}

}
