package main

import (
	"fmt"
	"os"
	"io"
	)

func main(){
	f, err := os.Open("messages.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer f.Close()
	
	buffer := make([]byte, 8)

	for{
		_, err := f.Read(buffer)
		if err == io.EOF{
			fmt.Println(err)
			break
		}
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Printf("%s \n",buffer[:n])
	}
}
