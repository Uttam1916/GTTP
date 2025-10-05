package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string{
	lines := make(chan string)

	go func(){

		defer close(lines)
		defer f.Close()

		buffer := make([]byte, 8)
		currentLine := ""
		
		for {

			n, err := f.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error reading file:", err)
				break
			}

			chunks := string(buffer[:n])
			parts := strings.Split(chunks,"\n")

			for i :=0 ; i < len(parts)-1 ; i++ {
				lines <- currentLine+parts[i]
				currentLine = ""
			}

			currentLine += parts[len(parts)-1]
		}:

		if currentLine != "" {
			lines <- currentLine
		}

	}()

	return lines
}

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for line := range getLinesChannel(f){
		fmt.Printf("read: %s \n",line)
	}
}
