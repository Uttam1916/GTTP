package main

import (
	"fmt"
	"io"
	"strings"
	"net"
	"log"
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
		}

		if currentLine != "" {
			lines <- currentLine
		}

	}()

	return lines
}

func main() {

	ln, err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("listening on port 8080")
	defer ln.Close()
	
	for{
		con, err := ln.Accept()
		if err!=nil{
			log.Println(err)
			continue
		}
		fmt.Println("connection established")

		for line := range getLinesChannel(con){
			fmt.Printf("read: %s \n",line)
		}

		fmt.Println("Connection ended")
	}
}
