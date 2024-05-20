package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fileSize, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	//fileSize, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", fileSize)

	f.Close()

	//leitura
	file, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	//leitura de arquivo em partes
	readFileInParts, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(readFileInParts)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println((string(buffer[:n])))
	}
	
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
